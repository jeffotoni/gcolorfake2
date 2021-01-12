package gcolorfake2 

const Sufix = "\033[0m"

func Yellow(str string) string {
   return "\033[0;33m" + str+ Sufix
}
