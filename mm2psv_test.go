package main

import "testing"
import "os"
import (
	"io"
	"bytes"
	"strings"
)

var testXml = `<map version="1.0.1">
<!-- To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net -->
<node CREATED="1445913736403" ID="ID_1735430275" MODIFIED="1445914098478" TEXT="fish">
<node CREATED="1445913745251" ID="ID_1555435469" MODIFIED="1445913777031" POSITION="right" TEXT="jawless">
<node CREATED="1445913778179" ID="ID_515153807" MODIFIED="1445913781606" TEXT="lampreys"/>
<node CREATED="1445913782096" ID="ID_913182564" MODIFIED="1445913784910" TEXT="hagfish"/>
</node>
<node CREATED="1445913763588" ID="ID_574727681" MODIFIED="1445914053825" POSITION="left" TEXT="cartilanginous">
<node CREATED="1445913793763" ID="ID_879169234" MODIFIED="1445913798469" TEXT="sharks"/>
<node CREATED="1445913799142" ID="ID_1181545561" MODIFIED="1445913818261" TEXT="rays"/>
<node CREATED="1445913818782" ID="ID_1884072293" MODIFIED="1445913821781" TEXT="chimaera"/>
</node>
<node CREATED="1445913825107" ID="ID_1314168113" MODIFIED="1445913827900" POSITION="left" TEXT="bony">
<node CREATED="1445913830033" ID="ID_1839458394" MODIFIED="1445913833212" TEXT="lobe finned">
<node CREATED="1445913840917" ID="ID_1153905153" MODIFIED="1445913843475" TEXT="lungfish"/>
<node CREATED="1445913876832" ID="ID_65064644" MODIFIED="1445913881465" TEXT="coelacanths"/>
</node>
<node CREATED="1445913835617" ID="ID_794440360" MODIFIED="1445913839099" TEXT="ray finned">
<node CREATED="1445913857011" ID="ID_1302114673" MODIFIED="1445913868890" TEXT="chondrosteans"/>
<node CREATED="1445913869619" ID="ID_478665094" MODIFIED="1445913873946" TEXT="holosteans"/>
<node CREATED="1445913883121" ID="ID_624030366" MODIFIED="1445913886017" TEXT="teleosts"/>
</node>
</node>
</node>
</map>`

func TestProcessXml(t *testing.T) {
	orig := os.Stdout // backup of real stdout
	r,w,_ := os.Pipe()
	os.Stdout = w


	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	processXmlString([]byte(testXml))

	w.Close()
	os.Stdout = orig // restore stdout
	out := <-outC

	if (strings.Count(out, "|fish|") != 10) {
		t.Error("Failure! Not enough rows.")
	}
	if (strings.Contains(out,"|fish|bony|lobe finned|lungfish|") == false) {
		t.Error("Failure! Missing lungfish in test.")
	}
}


