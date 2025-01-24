package formatter

import (
	"encoding/json"
)

func uniqueMedia(msg *Message) error {
	var err error
	var mediajson []byte
	if msg.fm.mh.atLeastOnce {
		if err = msg.fm.mh.storage[msg.fm.mh.i-1].multipartFields(msg.fm.writer, nil, 0, false); err == nil {
			msg.fm.contentType = msg.fm.writer.FormDataContentType()
		}
	} else {
		if mediajson, err = msg.fm.mh.storage[msg.fm.mh.i-1].jsonFileds(); err == nil {
			msg.fm.buf.Write(mediajson)
		}
	}
	return err
}

func mediaGroup(msg *Message) error {
	var (
		err    error
		jsbody []byte
	)
	group := make([]interface{}, msg.fm.mh.amount)
	mediaMap := make(map[string]interface{}, 1)
	if !msg.fm.mh.atLeastOnce {
		for i := 0; i < len(msg.fm.mh.storage); i++ {

			if msg.fm.mh.storage[i] != nil {
				group[i] = msg.fm.mh.storage[i]
			}
		}

		mediaMap["media"] = group
		jsbody, err = json.Marshal(mediaMap)
		if err == nil {
			msg.fm.buf.Write(jsbody)
		}
	} else {

		for i := 0; i < len(msg.fm.mh.storage) && err == nil; i++ {
			if msg.fm.mh.storage[i] != nil {
				err = msg.fm.mh.storage[i].multipartFields(msg.fm.writer, &group, i, true)
			}
		}
		if err == nil {
			err = putGroup(msg.fm.writer, group)
		}
		msg.fm.contentType = msg.fm.writer.FormDataContentType()
	}
	return err
}
