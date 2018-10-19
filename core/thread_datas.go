package core

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/textileio/textile-go/ipfs"
	"github.com/textileio/textile-go/pb"
	"github.com/textileio/textile-go/repo"
	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
	"gx/ipfs/QmdVrMn1LhB4ybb8hMVaMLXnA8XRSewMnK6YqXKXoTcRvN/go-libp2p-peer"
)

// AddPhoto adds an outgoing photo block
func (t *Thread) AddPhoto(dataId string, caption string, key []byte) (mh.Multihash, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	// download metadata
	metadataCipher, err := ipfs.GetDataAtPath(t.node(), fmt.Sprintf("%s/meta", dataId))
	if err != nil {
		return nil, err
	}

	// encrypt AES key with thread pk
	keyCipher, err := t.Encrypt(key)
	if err != nil {
		return nil, err
	}

	// encrypt caption with thread pk
	var captionCipher []byte
	if caption != "" {
		captionCipher, err = t.Encrypt([]byte(caption))
		if err != nil {
			return nil, err
		}
	}

	// build block
	header, err := t.newBlockHeader()
	if err != nil {
		return nil, err
	}
	content := &pb.ThreadData{
		Header:        header,
		Type:          pb.ThreadData_PHOTO,
		DataId:        dataId,
		KeyCipher:     keyCipher,
		CaptionCipher: captionCipher,
	}

	// commit to ipfs
	env, addr, err := t.commitBlock(content, pb.Message_THREAD_DATA)
	if err != nil {
		return nil, err
	}
	id := addr.B58String()

	// index it locally
	dconf := &repo.DataBlockConfig{
		DataId:             dataId,
		DataKeyCipher:      keyCipher,
		DataCaptionCipher:  captionCipher,
		DataMetadataCipher: metadataCipher,
	}
	if err := t.indexBlock(id, header, repo.PhotoBlock, dconf); err != nil {
		return nil, err
	}

	// update head
	if err := t.updateHead(id); err != nil {
		return nil, err
	}

	// post it
	t.post(env, id, t.Peers())

	log.Debugf("added DATA to %s: %s", t.Id, id)

	// all done
	return addr, nil
}

// HandleDataBlock handles an incoming data block
func (t *Thread) HandleDataBlock(from *peer.ID, env *pb.Envelope, signed *pb.SignedThreadBlock, content *pb.ThreadData, following bool) (mh.Multihash, error) {
	// unmarshal if needed
	if content == nil {
		content = new(pb.ThreadData)
		if err := proto.Unmarshal(signed.Block, content); err != nil {
			return nil, err
		}
	}

	// add to ipfs
	addr, err := t.addBlock(env)
	if err != nil {
		return nil, err
	}
	id := addr.B58String()

	// check if we aleady have this block indexed
	// (should only happen if a misbehaving peer keeps sending the same block)
	index := t.datastore.Blocks().Get(id)
	if index != nil {
		return nil, nil
	}

	// get the author id
	authorId, err := ipfs.IDFromPublicKeyBytes(content.Header.AuthorPk)
	if err != nil {
		return nil, err
	}

	// add author as a new local peer, just in case we haven't found this peer yet.
	// double-check not self in case we're re-discovering the thread
	self := authorId.Pretty() == t.node().Identity.Pretty()
	if !self {
		threadId, err := ipfs.IDFromPublicKeyBytes(content.Header.ThreadPk)
		if err != nil {
			return nil, err
		}
		newPeer := &repo.ThreadPeer{
			Id:       authorId.Pretty(),
			ThreadId: threadId.Pretty(),
		}
		if err := t.datastore.ThreadPeers().Add(newPeer); err != nil {
			log.Errorf("error adding peer: %s", err)
		}
	}

	// index it locally
	dconf := &repo.DataBlockConfig{
		DataId:            content.DataId,
		DataKeyCipher:     content.KeyCipher,
		DataCaptionCipher: content.CaptionCipher,
	}
	switch content.Type {
	case pb.ThreadData_PHOTO:
		// check if this block has been ignored, if so, don't pin locally, but we still want the block
		var ignore bool
		ignored := t.datastore.Blocks().GetByData(fmt.Sprintf("ignore-%s", id))
		if ignored != nil {
			date, err := ptypes.Timestamp(content.Header.Date)
			if err != nil {
				return nil, err
			}
			// ignore if the ignore block came after (could happen when bootstrapping the thread during back prop)
			if ignored.Date.After(date) {
				ignore = true
			}
		}
		if !ignore {
			// pin data first (it may not be available)
			if err := ipfs.PinPath(t.node(), fmt.Sprintf("%s/thumb", content.DataId), false); err != nil {
				return nil, err
			}
			if err := ipfs.PinPath(t.node(), fmt.Sprintf("%s/small", content.DataId), false); err != nil {
				log.Warningf("photo set missing small size")
			}
			if err := ipfs.PinPath(t.node(), fmt.Sprintf("%s/meta", content.DataId), false); err != nil {
				return nil, err
			}
			if err := ipfs.PinPath(t.node(), fmt.Sprintf("%s/pk", content.DataId), false); err != nil {
				return nil, err
			}
		}

		// get metadata
		metadataCipher, err := ipfs.GetDataAtPath(t.node(), fmt.Sprintf("%s/meta", content.DataId))
		if err != nil {
			return nil, err
		}
		dconf.DataMetadataCipher = metadataCipher

		// index
		if err := t.indexBlock(id, content.Header, repo.PhotoBlock, dconf); err != nil {
			return nil, err
		}
	case pb.ThreadData_TEXT:
		// TODO: chat
		break
	}

	// back prop
	newPeers, err := t.FollowParents(content.Header.Parents, from)
	if err != nil {
		return nil, err
	}

	// handle HEAD
	if following {
		return addr, nil
	}
	if _, err := t.handleHead(id, content.Header.Parents); err != nil {
		return nil, err
	}

	// handle newly discovered peers during back prop, after updating HEAD
	for _, newPeer := range newPeers {
		if err := t.sendWelcome(newPeer); err != nil {
			return nil, err
		}
	}

	return addr, nil
}
