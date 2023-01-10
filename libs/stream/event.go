package stream

import (
	"log"
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

type event struct {
	streamName string
	subject    string
}

// Parses a protobuf message to an event.
func eventFromProtobufMessage(proto any) event {
	parts := protoToParts(proto)

	streamName := partsToStreamName(parts)
	subjectSuffix := partsToSubjectSuffix(parts)

	return event{
		streamName: streamName,
		subject:    streamName + "." + subjectSuffix,
	}
}

var regexMatchUpperCase = regexp.MustCompile(`[A-Z][^A-Z]*`)

// This returns the raw stream name and suffix of the subject for a protbuf event.
// To publish the event we still need to format the two parts
// and concat the suffix of the subject with the stream name to get the publishable subject.
func protoToParts(event any) []string {
	t := reflect.TypeOf(event)
	str := strings.ReplaceAll(t.String(), "*", "")

	split := strings.Split(str, ".")

	withoutPackage := str

	// the event is probably from a different package so it has the package prefix which we can ignore
	if len(split) > 1 {
		withoutPackage = strings.Join(split[1:], ".")
	}

	parts := regexMatchUpperCase.FindAllString(withoutPackage, -1)

	if parts == nil || len(parts) < 2 {
		log.Fatalln("eventToStreamName: Parsed parts of event are invalid: ", parts)
	}

	// this allows subject suffix like CreatedNow
	parts[1] = strings.Join(parts[1:], "")

	// splits the event name into the stream name and suffix of subject
	return parts
}

func partsToStreamName(parts []string) string {
	return strings.ToUpper(parts[0])
}

func partsToSubjectSuffix(parts []string) string {
	return camelcaseStringToDotString(parts[1])
}

func camelcaseStringToDotString(camelcase string) string {
	var b strings.Builder

	for i, c := range camelcase {
		if unicode.IsUpper(c) {
			if i != 0 {
				b.WriteString(".")
			}
			b.WriteRune(unicode.ToLower(c))
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}
