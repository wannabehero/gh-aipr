package utils

import "math/rand"

var emojis = []string{
	"🚀", "🤖", "🎢", "🪨", "🎨",
	"🧻", "🗞️", "📎", "🌚", "💸",
	"📝", "☠️", "🐕", "🐩", "🐃",
	"♾️", "🐜", "🦁", "🐺", "🦊",
	"🦁", "🐯", "🐘", "🐍", "🦄",
	"🔥", "💫", "⚡", "🌈", "🍕",
	"🧠", "👾", "🤖", "🎮", "🚗",
	"🌳", "🌵", "🍄", "🦅", "🦉",
}

func GetRandomEmoji() string {
	return emojis[rand.Intn(len(emojis))]
}
