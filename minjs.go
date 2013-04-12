package gomin

import (
	"errors"
)

var err error

var input []byte
var inputIndex int = 0
var output []byte

var theA byte
var theB byte
var theLookahead byte = 0
var theX byte = 0
var theY byte = 0

func isalnum(c byte) bool {
	return '0' <= c && c <= '9' || 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z'
}

func get() byte {
	var c byte = theLookahead
	theLookahead = 0
	if c == 0 {
		if inputIndex < len(input) {
			c = input[inputIndex]
			inputIndex += 1
		} else {
			c = 0
		}
	}
	if c >= ' ' || c == '\n' || c == 0 {
		return c
	}
	if c == '\r' {
		return '\n'
	}
	return ' '
}

func peek() byte {
	theLookahead = get()
	return theLookahead
}

func next() (byte, error) {
	var c byte = get()
	if c == '/' {
		switch peek() {
		case '/':
			for {
				c = get()
				if c <= '\n' {
					break
				}
			}
		case '*':
			get()
			for c != ' ' {
				switch get() {
				case '*':
					if peek() == '/' {
						get()
						c = ' '
					}
				case 0:
					return 0, errors.New("Unterminated comment.")
				}
			}
		}
	}
	theY = theX
	theX = c
	return c, nil
}

func action(d int) error {
	if d == 1 {
		output = append(output, theA)
		if (theY == '\n' || theY == ' ') &&
			(theA == '+' || theA == '-' || theA == '*' || theA == '/') &&
			(theB == '+' || theB == '-' || theB == '*' || theB == '/') {
			output = append(output, theY)
		}
	}
	if d <= 2 {
		theA = theB
		if theA == '\'' || theA == '"' || theA == '`' {
			for {
				output = append(output, theA)
				theA = get()
				if theA == theB {
					break
				}
				if theA == '\\' {
					output = append(output, theA)
					theA = get()
				}
				if theA == 0 {
					return errors.New("Unterminated string literal.")
				}
			}
		}
	}
	if d <= 3 {
		theB, err = next()
		if err != nil {
			return err
		}
		if theB == '/' &&
			(theA == '(' || theA == ',' || theA == '=' || theA == ':' ||
				theA == '[' || theA == '!' || theA == '&' || theA == '|' ||
				theA == '?' || theA == '+' || theA == '-' || theA == '~' ||
				theA == '*' || theA == '/' || theA == '{' || theA == '\n') {
			output = append(output, theA)
			if theA == '/' || theA == '*' {
				output = append(output, ' ')
			}
			output = append(output, theB)
			for {
				theA = get()
				if theA == '[' {
					for {
						output = append(output, theA)
						theA = get()
						if theA == ']' {
							break
						}
						if theA == '\\' {
							output = append(output, theA)
							theA = get()
						}
						if theA == 0 {
							return errors.New("Unterminated set in Regular Expression literal.")
						}
					}
				} else if theA == '/' {
					switch peek() {
					case '/', '*':
						return errors.New("Unterminated set in Regular Expression literal.")
					}
					break
				} else if theA == '\\' {
					output = append(output, theA)
					theA = get()
				}
				if theA == 0 {
					return errors.New("Unterminated set in Regular Expression literal.")
				}
				output = append(output, theA)
			}
			theB, err = next()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func MinJS(s []byte) ([]byte, error) {
	input = s
	if peek() == 0xEF {
		get()
		get()
		get()
	}
	theA = '\n'
	err = action(3)
	if err != nil {
		return output, err
	}
	for theA != 0 {
		switch theA {
		case ' ':
			if isalnum(theB) {
				err = action(1)
			} else {
				err = action(2)
			}
		case '\n':
			switch theB {
			case '{', '[', '(', '+', '-', '!', '~':
				err = action(1)
			case ' ':
				err = action(3)
			default:
				if isalnum(theB) {
					err = action(1)
				} else {
					err = action(2)
				}
			}
		default:
			switch theB {
			case ' ':
				if isalnum(theA) {
					err = action(1)
				} else {
					err = action(3)
				}
			case '\n':
				switch theA {
				case '}', ']', ')', '+', '-', '"', '\'', '`':
					err = action(1)
				default:
					if isalnum(theA) {
						err = action(1)
					} else {
						err = action(3)
					}
				}
			default:
				err = action(1)
			}
		}
		if err != nil {
			return output, err
		}
	}
	return output, nil
}
