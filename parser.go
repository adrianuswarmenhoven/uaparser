package uaparser

import (
    "strings"
    "unicode"
)

type itemSpec struct {
    name string
    mustContains []string
    mustNotContains []string
    versionSplitters [][]string
}

type InfoItem struct {
    Name string
    Version string
}

type UAInfo struct {
    Browser,
    Device,
    OS *InfoItem
}

func isEmptyString(str string) bool {
    for _, char := range str {
        if !unicode.IsSpace(char) {
            return false
        }
    }
    return true
}

func matchSpec(ua string, spec *itemSpec) (info *InfoItem, ok bool) {
    for _, mc := range spec.mustContains {
        if !strings.Contains(ua, mc) {
            return
        }
    }

    for _, mnc := range spec.mustNotContains {
        if strings.Contains(ua, mnc) {
            return
        }
    }

    info = new(InfoItem)
    info.Name = spec.name
    ok = true

    for _, splitter := range spec.versionSplitters {
        if strings.Contains(ua, splitter[0]) {
            if rmLeft := strings.Split(ua, splitter[0])[1];
               strings.Contains(rmLeft, splitter[1]) || isEmptyString(splitter[1]) {
                rmRight := strings.Split(rmLeft, splitter[1])[0]
                info.Version = strings.TrimSpace(rmRight)
                break
            }
        }
    }
    return
}

func searchIn(ua string, specs []*itemSpec) (info *InfoItem) {
    for _, spec := range specs {
        if result, ok := matchSpec(ua, spec); ok {
            info = result
            break
        }
    }
    return
}

func Parse(ua string) (info *UAInfo) {
    info = new(UAInfo)

    info.Browser = searchIn(ua, _BROWSERS)
    info.Device = searchIn(ua, _DEVICES)
    info.OS = searchIn(ua, _OS)

    return
}

