package helpers

import "fmt"

func GenerateAvatarURL(email, role string) string {
    set := "set4"
    if role == "trainer" {
        set = "set2"
    }
    return fmt.Sprintf("https://robohash.org/%s?set=%s", email, set)
}
