// gpx-fix-rollover: Fixes the GPS week rollover issue in GPC files
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
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/tkrajina/gpxgo/gpx"
)

const defaultOffset = "172032h" // 1024 weeks

func main() {
	fmt.Println("gpx-fix-rollover: Fixes the GPS week rollover issue in GPX files.")
	fmt.Println("Licensed under the GNU GPL version 3 or later.\n")

	// Command-line arguments
	timeOffsetFlag := flag.String("offset", defaultOffset, "The time offset to add to add to each point. Accepts strings like '12h3m15s'. Default is 1024 weeks.")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s <input-file.gpx> [flags]\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Available flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		return
	}

	timeOffset, err := time.ParseDuration(*timeOffsetFlag)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing duration %s: %s\n", timeOffsetFlag, err.Error())
		return
	}

	inFile := flag.Arg(0)
	backupFile := inFile + ".backup"
	outFile := inFile

	if _, err := os.Stat(backupFile); err == nil {
		fmt.Fprintf(os.Stderr, "Backup file %s already exists.\n", backupFile)
		return
	} else if !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Cannot create backup file %s: %s\n", backupFile, err.Error())
		return
	}

	// Do the job!
	fmt.Println("Working...")

	gpxFile, err := gpx.ParseFile(inFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file %s: %s\n", inFile, err.Error())
		return
	}

	gpxFile.ExecuteOnAllPoints(func(p *gpx.GPXPoint) { p.Timestamp = p.Timestamp.Add(timeOffset) })

	err = os.Rename(inFile, backupFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error backing the GPX file up: %s\n", err.Error())
		return
	}

	xmlBytes, err := gpxFile.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating output GPX: %s\n", err.Error())
		return
	}

	err = ioutil.WriteFile(outFile, xmlBytes, 0644)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to output file %s: %s\n", outFile, err.Error())
		return
	}

	// All done! Be nice end exit.
	fmt.Println("Done! Have a nice day!")
}
