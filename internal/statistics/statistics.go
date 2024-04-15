package statistics

import (
	"bytes"

	"fmt"

	"github.com/jsightapi/datagram"

	"github.com/jsightapi/jsight-api-core/jerr"

	"github.com/jsightapi/jsight-api-core/kit"

	"log"

	"net"

	"strconv"

	"github.com/denisbrodbeck/machineid"
)

var serverStatAddress = "127.0.0.1:1053" // "stat.jsight.io:1053"

const macAddressLength = 17

func SendStat(j *kit.JApi, je *jerr.JApiError, sendStatFlag bool, fileSize int64, err error) {
	if sendStatFlag {
		id, iderr := machineid.ID()
		if iderr != nil {
			id = UUIDMacID() // return uuid from the first of mac addresses
		}

		if err != nil && je == nil {
			ee := jerr.JApiError{}
			ee.Msg = err.Error()
			ee.Index = 0
			ee.Line = 0
			je = &ee
		}

		sendDatagram(id, "", int(fileSize), j, je)
	}
}

func sendDatagram(clientID, clientIP string, projectSize int, j *kit.JApi, je *jerr.JApiError) {
	title := ""
	if j != nil {
		title = j.Title()
	}
	d := datagram.New()
	d.Append("cid", clientID)
	d.Append("cip", clientIP)
	d.Append("at", "2")                       // Application Type, 2 for CLI, constant
	d.AppendTruncatable("pt", title)          // Project title
	d.Append("ps", strconv.Itoa(projectSize)) // Project size
	if je != nil {
		d.AppendTruncatable("pem", je.Error())                    // Project error message
		d.Append("pel", strconv.FormatUint(uint64(je.Line), 10))  // Project error line
		d.Append("pei", strconv.FormatUint(uint64(je.Index), 10)) // Project error index
	}

	err := sendToStatisticServer(d.Pack())
	if err != nil {
		log.Print("... " + err.Error())
	}
}

func sendToStatisticServer(b []byte) error {
	a, err := net.ResolveUDPAddr("udp4", serverStatAddress)
	if err != nil {
		return err
	}

	c, err := net.DialUDP("udp4", nil, a)
	if err != nil {
		return err
	}

	defer func() {
		_ = c.Close()
	}()

	_, err = c.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func UUIDMacID() string {
	as, err := macAddrList()
	if err == nil {
		for _, a := range as {
			if len(a) == macAddressLength {
				return fmt.Sprintf("00000000-0000-0000-0000-%s%s%s%s%s%s", a[0:2], a[3:5], a[6:8], a[9:11], a[12:14], a[15:17])
			}
		}
	}
	return "0"
}

func macAddrList() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		if ifa.Flags&net.FlagUp != 0 && !bytes.Equal(ifa.HardwareAddr, nil) {
			a := ifa.HardwareAddr.String()
			if a != "" {
				as = append(as, a)
			}
		}
	}
	return as, nil
}
