// gpx_fix_rollover: Fixes the GPS week rollover issue in GPC files
//
// Copyright (C) 2019  Leandro Motta Barros
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
)

func main() {
	fmt.Println("gpx_fix_rollover: Fixes the GPS week rollover issue in GPX files.")
	fmt.Println("Licensed under the GNU GPL version 3 or later.")
	fmt.Println("Working...")

	const inFile = "20190528-191450.gpx"
	const outFile = "result.gpx"
	const deltaTimeDurationString = "172032h" // 1024 weeks

	deltaTime, err := time.ParseDuration(deltaTimeDurationString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing duration %s: %s", deltaTimeDurationString, err.Error())
		return
	}

	gpxBytes, err := ioutil.ReadFile(inFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading file %s: %s", inFile, err.Error())
		return
	}

	// TODO: use ParseFile
	gpxFile, err := gpx.ParseBytes(gpxBytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file %s: %s", inFile, err.Error())
		return
	}

	gpxFile.ExecuteOnAllPoints(func(p *gpx.GPXPoint) { p.Timestamp = p.Timestamp.Add(deltaTime) })

	// When ready, you can write the resulting GPX file:
	xmlBytes, err := gpxFile.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file %s: %s", outFile, err.Error())
		return
	}

	ioutil.WriteFile(outFile, xmlBytes, 0644)

	fmt.Println("Done! Have a nice day!")
}
