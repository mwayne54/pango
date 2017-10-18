package addrgrp

import (
    "testing"
    "reflect"

    "github.com/PaloAltoNetworks/xapi/testdata"
)


func TestNormalization(t *testing.T) {
    testCases := []struct{
        desc string
        vsys string
        conf Entry
    }{
        {"test static no tags", "", Entry{
            Name: "one",
            Description: "my description",
            Static: []string{"adr1", "adr2"},
        }},
        {"test static with tags", "", Entry{
            Name: "one",
            Description: "my description",
            Static: []string{"adr1", "adr2"},
            Tag: []string{"tag1", "tag2"},
        }},
        {"test dynamic no tags", "", Entry{
            Name: "one",
            Description: "my description",
            Dynamic: "'tag1' or 'tag2' and 'tag3'",
        }},
        {"test dynamic with tags", "", Entry{
            Name: "one",
            Description: "my description",
            Dynamic: "'tag1' or 'tag2' and 'tag3'",
            Tag: []string{"tag1", "tag2"},
        }},
    }

    mc := &testdata.MockClient{}
    ns := &AddrGrp{}
    ns.Initialize(mc)

    for _, tc := range testCases {
        t.Run(tc.desc, func(t *testing.T) {
            mc.Reset()
            mc.AddResp("")
            err := ns.Set(tc.vsys, tc.conf)
            if err != nil {
                t.Errorf("Error in set: %s", err)
            } else {
                mc.AddResp(mc.Elm)
                r, err := ns.Get(tc.vsys, tc.conf.Name)
                if err != nil {
                    t.Errorf("Error in get: %s", err)
                } else if !reflect.DeepEqual(tc.conf, r) {
                    t.Errorf("%#v != %#v", tc.conf, r)
                }
            }
        })
    }
}

