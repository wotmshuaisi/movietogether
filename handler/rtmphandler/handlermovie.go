package rtmphandler

import (
	"time"

	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/format/rtmp"
	"github.com/sirupsen/logrus"
	"github.com/wotmshuaisi/movietogether/config"
)

// Publish accept rtmp streaming
func (handlers *RTMPHandler) Publish(rtmpCon *rtmp.Conn) {
	logrus.WithFields(logrus.Fields{
		"datetime": time.Now().String(),
		"host":     rtmpCon.URL.Host,
		"query":    rtmpCon.URL.RawQuery,
	}).Info()
	if rtmpCon.URL.Path != config.RTMPURL {
		rtmpCon.Close()
		return
	}
	// auth
	if rtmpCon.URL.Query().Get("token") != config.PUBLISHSECRET {
		rtmpCon.Close()
		return
	}
	// publish part
	stream, err := rtmpCon.Streams()
	if err != nil {
		logrus.WithError(err).Errorln("rtmp get streaming")
		return
	}

	handlers.Channel.WriteHeader(stream)
	avutil.CopyPackets(handlers.Channel, rtmpCon)
	return
}

// // Play accept rtmp streaming
// func (handlers *RTMPHandler) Play(rtmpCon *rtmp.Conn) {
// 	if handlers.Channel != nil {
// 		cursor := handlers.Channel.Latest()
// 		avutil.CopyFile(rtmpCon, cursor)
// 	}
// 	return
// }
