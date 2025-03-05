package integrated

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/l1qwie/Fmtogram/formatter"
	"github.com/l1qwie/Fmtogram/testbotdata"
	"github.com/l1qwie/Fmtogram/types"
)

const (
	chatid                                                int     = 738070596
	photo, video, audio, document, together               int     = 0, 1, 2, 3, 4
	link                                                  string  = "https://t.me/+azuTu6sZ5CBjNzA6"
	textformsg                                            string  = "There's some kind of text"
	status, errr, chat, sender, date, msgid               int     = 0, 1, 2, 3, 4, 5
	msgids, replyed, fororig, ph, ad, doc                 int     = 6, 7, 8, 9, 10, 11
	vd, anim, vc, vdn, paidm, groupid, phs                int     = 12, 13, 14, 15, 16, 17, 18
	vds, ads, docs, poll, dice, profph, file              int     = 19, 20, 21, 22, 23, 24, 25
	stkrs, gifts, msg, str, invl, chinf, memb             int     = 26, 27, 28, 29, 30, 31, 32
	integer, frm, boosts, bcon, comm, mbut, adright, msgs int     = 33, 34, 35, 36, 37, 38, 39, 40
	msgEffect                                             string  = "5107584321108051014"
	testybotid                                            int     = 8033103339
	unknownid                                             int     = -998425295
	supergroupid                                          int     = -1002309471573
	month                                                 int     = 2592000
	iconemojiid                                           string  = "C5310262535021142850"
	iconcolor                                             int     = 0x6FB9F0
	english                                               string  = "en"
	stars                                                 int     = 33
	latitude, longitude                                   float64 = 31.7962236, 35.0194702
	placetitle                                            string  = "The City Of The History"
	city                                                  string  = "Jerusalem"
	phonenum                                              string  = "0123422211"
	name                                                  string  = "Nathan"
	lastname                                              string  = "Sahar"
	vcard                                                 string  = "BEGIN:VCARD\nVERSION:4.0\nFN:Nathan Sahar\nN:Nathan;Simon;;;ing. jr,M.Sc.\nBDAY:--0203\nGENDER:M\nEMAIL;TYPE=work:nathan.sahar@proton.me\nEND:VCARD"
	question                                              string  = "Am I good?"
	explanation                                           string  = "Yes, I am!"
	thumbnail                                             string  = "../media/tel-aviv.jpg"
	photopath                                             string  = "../media/tel-aviv.jpg"
	photourl                                              string  = "https://www.aljazeera.com/wp-content/uploads/2025/01/AFP__20250120__36UX28A__v2__HighRes__TopshotUsPoliticsTrumpInauguration-1737420954.jpg?w=770&resize=770,513"
	videopath                                             string  = "../media/black.mp4"
	videourl                                              string  = "https://www.pexels.com/download/video/6646588/"
	audiopath                                             string  = "../media/sound.mp3"
	audiourl                                              string  = "https://freetestdata.com/wp-content/uploads/2021/09/Free_Test_Data_100KB_MP3.mp3"
	docpath                                               string  = "../media/Resume.pdf"
	docurl                                                string  = "https://www.aljazeera.com/wp-content/uploads/2025/01/AFP__20250120__36UX28A__v2__HighRes__TopshotUsPoliticsTrumpInauguration-1737420954.jpg?w=770&resize=770,513"
	animpath                                              string  = "../media/prichinatryski.mp4"
	animurl                                               string  = "https://www.icegif.com/wp-content/uploads/2025/02/patrick-mahomes-icegif-1.gif"
	voicepath                                             string  = "../media/dimaJOSKAproNATO.ogg"
	voiceurl                                              string  = "https://freetestdata.com/wp-content/uploads/2021/09/Free_Test_Data_100KB_OGG.ogg"
	vdnpath                                               string  = "../media/tea.mp4"
	vdnurl                                                string  = "https://www.pexels.com/download/video/6646588/"
	stickerpath                                           string  = "../media/sticker.webp"
	stickerurl                                            string  = "https://www.gstatic.com/webp/gallery/1.webp"
)

var (
	f bool
)

var photodata = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {photopath, "AgACAgIAAxkDAAIhPWe_Mzh3BuGZDMORZ2MJ45vudhYlAAL97TEb4pX5ScXXIqH4IsoOAQADAgADcwADNgQ", photourl},
	testbotdata.Tttesty_bot_bot:           {photopath, "AgACAgIAAxkDAAO6Z8B2X4OpA3NUy1xRl8TCNgHiGX0AAnzvMRt6LgFKfBjZfF0RJSwBAAMCAANzAAM2BA", photourl},
	testbotdata.TestsInGroup2_bot:         {photopath, "AgACAgIAAxkDAAMJZ8B5eVFTt2IROngxWyjLMBE0eqwAAhfzMRsHFQABSoI_0wdUeiz5AQADAgADcwADNgQ", photourl},
	testbotdata.TestsInGroup1_bot:         {photopath, "AgACAgIAAxkDAAMJZ8B5e5isR164Mgd6BxqkUaFnSOwAAsvqMRugtwFKG4b7w399LV0BAAMCAANzAAM2BA", photourl},
	testbotdata.Make_event_bot:            {photopath, "AgACAgIAAxkDAAMLZ8B8wNqF9PBsPqH5spLZX0Bno5kAAgPpMRv9kglKOOoUmiKYpkABAAMCAANzAAM2BA", photourl},
	testbotdata.Testdatbot_bot:            {photopath, "AgACAgIAAxkDAAIF22fAfL9PqhxUbNE6BNRlXicLzuNXAAJ27jEbvNUJSvw73_eYcWmJAQADAgADcwADNgQ", photourl},
	testbotdata.SpeakOnlyWithAnton_bot:    {photopath, "AgACAgIAAxkDAANNZ8B-bwKGMc1DWKQxlztfp87Rg5EAAoHyMRtypglKUvHXuuKYrNgBAAMCAANzAAM2BA", photourl},
	testbotdata.LearnSpanishOrEnglish_bot: {photopath, "AgACAgIAAxkDAAMaZ8B9CIPGATcm5VndMVvCY3LE8e0AAqPyMRvaAQhKvx04fqGb290BAAMCAANzAAM2BA", photourl},
	testbotdata.Testmy_bots_bot:           {photopath, "AgACAgIAAxkDAAMQZ8CPdrNu-H8Z8VlEohxKiNNw4YEAApLqMRt61AlKkv-JUGXyiW4BAAMCAANzAAM2BA", photourl},
}

var videodata = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {videopath, "BAACAgIAAxkDAAIkpWfA-ROTjJn2SgW3nn3b9hq4uSGZAALSYwACcV8ISkb19Qq6Dm-pNgQ", videourl},
	testbotdata.Tttesty_bot_bot:           {videopath, "BAACAgIAAxkDAAPWZ8D5cRFvuaQJp2O916HtKpCkZHIAAnt0AAJ6LglKifadYDyru4s2BA", videourl},
	testbotdata.TestsInGroup2_bot:         {videopath, "BAACAgIAAxkDAAMZZ8D5n_huO0afeiLF99NVuqSiLXIAAnZiAAIHFQhK0VfYv32lbhE2BA", videourl},
	testbotdata.TestsInGroup1_bot:         {videopath, "BAACAgIAAxkDAAMRZ8D5z-GOD-XUL0M_9TAev6xpjqIAAklkAAKgtwlKTHLfI0t9vhw2BA", videourl},
	testbotdata.Make_event_bot:            {videopath, "BAACAgIAAxkDAAMUZ8D57kKf0zvAj5Z47cwhNch6J6sAAhVzAAL9kglKdRlCvD8gbyc2BA", videourl},
	testbotdata.Testdatbot_bot:            {videopath, "BAACAgIAAxkDAAIF7WfA-g1aWdIbCyKwRd_6OCMpXlVuAAJyZwACvNUJSuU1fhbKbdKBNgQ", videourl},
	testbotdata.SpeakOnlyWithAnton_bot:    {videopath, "BAACAgIAAxkDAANXZ8D6TK2eoN5ZXOVJE0lKmCv-OdwAAshxAAJAxQhKx_Ard9EsfFA2BA", videourl},
	testbotdata.LearnSpanishOrEnglish_bot: {videopath, "BAACAgIAAxkDAAMlZ8D6fRJh5Kg2MIQ_aeyJhvmpn9cAAoNtAALaAQhKUElYBI6SFMI2BA", videourl},
	testbotdata.Testmy_bots_bot:           {videopath, "BAACAgIAAxkDAAMRZ8D6mF1Q4IeI0KTjqgGWhs8CP0YAAjFuAAJ61AlKJaPtnVpIUn42BA", videourl},
}

var audiodata = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {audiopath, "CQACAgIAAxkDAAIkrWfA_H_6X4Svb30G8mZ9fkFTlKD9AALVYwACcV8ISpjBU0JKuNplNgQ", audiourl},
	testbotdata.Tttesty_bot_bot:           {audiopath, "CQACAgIAAxkDAAPYZ8D8wFnEkxsh8UM5MKBFaWDqgoAAArJnAAK6YglKciuy33KZazA2BA", audiourl},
	testbotdata.TestsInGroup2_bot:         {audiopath, "CQACAgIAAxkDAAMaZ8D9G5D4XSpnjagz4nApTimzc2sAAnxiAAIHFQhKdJausotJFAc2BA", audiourl},
	testbotdata.TestsInGroup1_bot:         {audiopath, "CQACAgIAAxkDAAMTZ8D9NKnzEusuAVl3pYLGITubFYUAAkxkAAKgtwlKAQnfXL9Xo5k2BA", audiourl},
	testbotdata.Make_event_bot:            {audiopath, "CQACAgIAAxkDAAMXZ8D9Va8eubW9-4WK2-DfeluehjcAAiFzAAL9kglK4gOoUOoq1sM2BA", audiourl},
	testbotdata.Testdatbot_bot:            {audiopath, "CQACAgIAAxkDAAIF8GfA_Y4t2Qumz6yx2pB84lF9ea9HAAJ6ZwACvNUJSu9Rc24OT6JcNgQ", audiourl},
	testbotdata.SpeakOnlyWithAnton_bot:    {audiopath, "CQACAgIAAxkDAANZZ8D9rJlXfVrreTriQTa1iRN2FSMAAjxxAAJypglKuvWjXaTaQbE2BA", audiourl},
	testbotdata.LearnSpanishOrEnglish_bot: {audiopath, "CQACAgIAAxkDAAMnZ8D90HFzHKXz78cVeh84DKkLIKEAAoxtAALaAQhKz6rt_LFujqQ2BA", audiourl},
	testbotdata.Testmy_bots_bot:           {audiopath, "CQACAgIAAxkDAAMSZ8D97cElgjU6Jplt-c7ld2R72E8AAqVyAAJv4ghKGvCvaxnmWSk2BA", audiourl},
}

var docdata = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {docpath, "BQACAgIAAxkDAAIkr2fA_22SFdTq6qGbJCnq6SI-oFB7AALYYwACcV8ISl4lLV1mUZrSNgQ", docurl},
	testbotdata.Tttesty_bot_bot:           {docpath, "BQACAgIAAxkDAAPZZ8D_nupoVKpWqpOZk_k4OZkfFzQAArZnAAK6YglK7HsSO94DFOs2BA", docurl},
	testbotdata.TestsInGroup2_bot:         {docpath, "BQACAgIAAxkDAAMbZ8D_4FAiQUTgfgw59oRWs0n3y44AAoBiAAIHFQhK6HfyMl7hbec2BA", docurl},
	testbotdata.TestsInGroup1_bot:         {docpath, "BQACAgIAAxkDAAMUZ8EAAaJ0n9WAy6Kcj38gqgph77oCAAJOZAACoLcJSvW9_qFv7-wgNgQ", docurl},
	testbotdata.Make_event_bot:            {docpath, "BQACAgIAAxkDAAMYZ8EAAcQrgGhhTM6OwQXF4SDvqzpyAAIocwAC_ZIJSudZcIoKXHv-NgQ", docurl},
	testbotdata.Testdatbot_bot:            {docpath, "BQACAgIAAxkDAAIF8mfBAAHkHkB1VObTzqAxjKmv5TyZ_gACfmcAArzVCUq4N4hkWD4DXDYE", docurl},
	testbotdata.SpeakOnlyWithAnton_bot:    {docpath, "BQACAgIAAxkDAANaZ8EBCY62D1BE4A96kbuQhD71x5MAAj9xAAJypglKuleictgeO_02BA", docurl},
	testbotdata.LearnSpanishOrEnglish_bot: {docpath, "BQACAgIAAxkDAAMoZ8EBMmDiSvvinGdpQ5RDQ0LT1ocAAh1zAAIEbghKEEtYOdQOLgI2BA", docurl},
	testbotdata.Testmy_bots_bot:           {docpath, "BQACAgIAAxkDAAMTZ8EBSUVJ70S1T84VzZmxxo7sNlAAAjVuAAJ61AlK9_9L1qOSjUs2BA", docurl},
}

var animdata = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {animpath, "CgACAgIAAxkDAAIks2fBAm0aPJMxXrbeT-lbvEYOylv2AALdYwACcV8ISjoro4fqSYWnNgQ", animurl},
	testbotdata.Tttesty_bot_bot:           {animpath, "CgACAgIAAxkDAAPcZ8ECuJyl_jHbLgfVj-Y9_CMZu90AAod0AAJ6LglKimlZ5DYeAy82BA", animurl},
	testbotdata.TestsInGroup2_bot:         {animpath, "CgACAgIAAxkDAAMdZ8EDNYKwikbsAAH_t3eNQGdY1FGYAAKFYgACBxUISr__FA-8bSBANgQ", animurl},
	testbotdata.TestsInGroup1_bot:         {animpath, "CgACAgIAAxkDAAMWZ8EDYDbNfmlAwNxm-SikTzBYmtEAAlVkAAKgtwlK7_Ncwboc6-E2BA", animurl},
	testbotdata.Make_event_bot:            {animpath, "CgACAgIAAxkDAAMZZ8EDjVYXkTedYuI83-KjT9JHzv0AAi1zAAL9kglKVoXYJEcpXOk2BA", animurl},
	testbotdata.Testdatbot_bot:            {animpath, "CgACAgIAAxkDAAIF82fBA7Ol3s5aEQABrY8w985YGRbUYQACiGcAArzVCUrhJ16Oj_VxzDYE", animurl},
	testbotdata.SpeakOnlyWithAnton_bot:    {animpath, "CgACAgIAAxkDAANbZ8ED9__BjHgZ-PHNHiqklR5rZswAAkBxAAJypglKAtpG-VQ6cNQ2BA", animurl},
	testbotdata.LearnSpanishOrEnglish_bot: {animpath, "CgACAgIAAxkDAAMpZ8EESHtxrEDdAnVRw-54zqe2AwkAApptAALaAQhKOsqtFLvV-qc2BA", animurl},
	testbotdata.Testmy_bots_bot:           {animpath, "CgACAgIAAxkDAAMUZ8EEiwfNeUHE39Bpmxr_Jo6vh4EAAjluAAJ61AlK0TIVEnZ98CI2BA", animurl},
}

var voicedata = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {voicepath, "AwACAgIAAxkDAAIktmfBBhU8CIpa_8tyfS6cbyA9QIPQAALiYwACcV8ISlHfoP-FSr2SNgQ", voiceurl},
	testbotdata.Tttesty_bot_bot:           {voicepath, "AwACAgIAAxkDAAPdZ8EGNK52eUTNDb0hSqBiSiMzF38AAo50AAJ6LglKDfznvitwxvw2BA", voiceurl},
	testbotdata.TestsInGroup2_bot:         {voicepath, "AwACAgIAAxkDAAMfZ8EGP1ZzHHrNd81FgWfQG21g2o8AAo1iAAIHFQhKf6lgRn_wYR02BA", voiceurl},
	testbotdata.TestsInGroup1_bot:         {voicepath, "AwACAgIAAxkDAAMXZ8EGUFHTCS1CzceNAAFCX5tUawxMAAJbZAACoLcJSok2Z4pHzDwcNgQ", voiceurl},
	testbotdata.Make_event_bot:            {voicepath, "AwACAgIAAxkDAAMbZ8EGYUl7WM75nzT7rSzm-IhHLyUAAjRzAAL9kglKwZMvOQz1pT42BA", voiceurl},
	testbotdata.Testdatbot_bot:            {voicepath, "AwACAgIAAxkDAAIF9GfBBm7uNKiCRONl-cceyxJic8RCAAKLZwACvNUJSqBoobsGztFRNgQ", voiceurl},
	testbotdata.SpeakOnlyWithAnton_bot:    {voicepath, "AwACAgIAAxkDAANfZ8EGgYDEv0Kiyn2h1quc-KLIoVQAAkVxAAJypglKDOb60JEaeek2BA", voiceurl},
	testbotdata.LearnSpanishOrEnglish_bot: {voicepath, "AwACAgIAAxkDAAMqZ8EGi6VhAAGOWpqLHzWZvGgAAQLz6gACMnMAAgRuCEpm_hZg2enGvTYE", voiceurl},
	testbotdata.Testmy_bots_bot:           {voicepath, "AwACAgIAAxkDAAMWZ8EGmmxIOE8-Njd4YrmUWfR-BWEAAj9uAAJ61AlKl_Ara4a7zbY2BA", voiceurl},
}

var videoNdata = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {vdnpath, "BAACAgIAAxkDAAIkvGfBkvLSP9hFU3NIc7vklpoPe7rvAAKHaAACcV8ISkiUHwhVsyfCNgQ", vdnurl},
	testbotdata.Tttesty_bot_bot:           {vdnpath, "BAACAgIAAxkDAAPgZ8GTjZ9cz0ItPSFej1DdSyZVmWsAAlBmAAJ6LhFK0QWpUbKmq902BA", vdnurl},
	testbotdata.TestsInGroup2_bot:         {vdnpath, "BAACAgIAAxkDAAMhZ8GTp9o5cP2NyD-OuMc-Z3Zk5xoAAuRmAAIHFQhKe7szag7Rhog2BA", vdnurl},
	testbotdata.TestsInGroup1_bot:         {vdnpath, "BAACAgIAAxkDAAMcZ8GT9OiHradGiCv3tLP69McwTUkAAkhpAAKgtwlKiwtDD_rXyWI2BA", vdnurl},
	testbotdata.Make_event_bot:            {vdnpath, "BAACAgIAAxkDAAMgZ8GUGMwxCk16CnPogd_nqoQCIr8AAjlgAAL9khFKE-qFFnndiWE2BA", vdnurl},
	testbotdata.Testdatbot_bot:            {vdnpath, "BAACAgIAAxkDAAIF9mfBlFlPVqlfnW-odJFCG23nnSXWAALoZgACvNURSoVcdVVbKFmaNgQ", vdnurl},
	testbotdata.SpeakOnlyWithAnton_bot:    {vdnpath, "BAACAgIAAxkDAANgZ8GUcuK2gaFSdO0PRgEIqKgrn6sAAkFfAAJyphFKlM67MZwJE-42BA", vdnurl},
	testbotdata.LearnSpanishOrEnglish_bot: {vdnpath, "BAACAgIAAxkDAAMsZ8GUmsK7YQ5TVZ0M9JYfyB6-NFUAAg9sAALaARBKEyhR3n97IEU2BA", vdnurl},
	testbotdata.Testmy_bots_bot:           {vdnpath, "BAACAgIAAxkDAAMYZ8GU1rUDG_jA0uCiJB5v7WufPcsAAsBbAAJ61BFKapM2KIL3jW82BA", vdnurl},
}

var stickerdata = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {stickerpath, "CAACAgIAAxkDAAIl5WfHo8zs5ud4xNaPxEmKXfPTJ5nDAAIxbAAC4rlBSpNZGOcIr2z6NgQ", stickerurl},
	testbotdata.Tttesty_bot_bot:           {stickerpath, "CAACAgIAAxkDAAIB32fHo-IWY706YU_LKIgkTlOHIl6kAAL9XQACsd9ASnlhL0Ds2pBONgQ", stickerurl},
	testbotdata.TestsInGroup2_bot:         {stickerpath, "CAACAgIAAxkDAAIBMmfHpBNEjb9N5bpau88H5wTOCat2AAIeXQACIr5ASklDvgmDNQcdNgQ", stickerurl},
	testbotdata.TestsInGroup1_bot:         {stickerpath, "CAACAgIAAxkDAAIBdGfHpCNQIP9ySZNvepLjdzWcR4TRAAKkdgAC9WY4Sp3x99p_5HslNgQ", stickerurl},
	testbotdata.Make_event_bot:            {stickerpath, "CAACAgIAAxkDAAIBIWfHpEOGChFq9qbh91ObA1rF0Q3NAAKgbQACD1FASlGLo2OH1YFrNgQ", stickerurl},
	testbotdata.Testdatbot_bot:            {stickerpath, "CAACAgIAAxkDAAIHPGfHpFDgAyKeicXGkhe1yUpcu_mCAAKqaQACO9hBSq6eJJQ8r6A2NgQ", stickerurl},
	testbotdata.SpeakOnlyWithAnton_bot:    {stickerpath, "CAACAgIAAxkDAAIBUWfHpF-tuT26e9zrC1Au-eZhYMDjAAJmaAACXjdBSiCPLuzh6PTJNgQ", stickerurl},
	testbotdata.LearnSpanishOrEnglish_bot: {stickerpath, "CAACAgIAAxkDAAIBTWfHpG0NF5SrRbEbqveHuO5v6_XyAAKGcQACNoA5SpvZY4yvTOMWNgQ", stickerurl},
	testbotdata.Testmy_bots_bot:           {stickerpath, "CAACAgIAAxkDAAMZZ8ekfON-YvQ9S5uw4IDx7j-zn9cAAut0AAJdwUFKndGRr9FWp8U2BA", stickerurl},
}

var thumbaudio = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {thumbnail, "AAMCAgADGQMAAiStZ8D8f_pfhK9vfQbyZn1-QVOUoP0AAtVjAAJxXwhKmMFTQkq42mUBAAdtAAM2BA"},
	testbotdata.Tttesty_bot_bot:           {thumbnail, "AAMCAgADGQMAA9hnwPzAWcSTGyHxQzkwoEVpYOqCgAACsmcAArpiCUpyK7LfcplrMAEAB20AAzYE"},
	testbotdata.TestsInGroup2_bot:         {thumbnail, "AAMCAgADGQMAAxpnwP0bkPhdKmeNqDPicClOKbNzawACfGIAAgcVCEp0lq6yi0kUBwEAB20AAzYE"},
	testbotdata.TestsInGroup1_bot:         {thumbnail, "AAMCAgADGQMAAxNnwP00qfMS6y4BWXelgsYhO5sVhQACTGQAAqC3CUoBCd9cv1ejmQEAB20AAzYE"},
	testbotdata.Make_event_bot:            {thumbnail, "AAMCAgADGQMAAxdnwP1Vrx65tb37hYrb4N96W56GNwACIXMAAv2SCUriA6hQ6irWwwEAB20AAzYE"},
	testbotdata.Testdatbot_bot:            {thumbnail, "AAMCAgADGQMAAgXwZ8D9ji3ZC6bPrLHakHziUX15r0cAAnpnAAK81QlK71Fzbg5PolwBAAdtAAM2BA"},
	testbotdata.SpeakOnlyWithAnton_bot:    {thumbnail, "AAMCAgADGQMAA1lnwP2smVd9Wut5OuJBNrWJE3YVIwACPHEAAnKmCUq69aNdpNpBsQEAB20AAzYE"},
	testbotdata.LearnSpanishOrEnglish_bot: {thumbnail, "AAMCAgADGQMAAydnwP3QcXMcpfPvxxV6HzgMqQsgoQACjG0AAtoBCErPqu38sW6OpAEAB20AAzYE"},
	testbotdata.Testmy_bots_bot:           {thumbnail, "AAMCAgADGQMAAxJnwP3twSWCNTommW35zuV3ZHvYTwACpXIAAm_iCEoa8K9rGeZZKQEAB20AAzYE"},
}

var thumbdoc = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {thumbnail, "AAMCAgADGQMAAiSvZ8D_bZIV1OrqoZskKerpIj6gUHsAAthjAAJxXwhKXiUtXWZRmtIBAAdtAAM2BA"},
	testbotdata.Tttesty_bot_bot:           {thumbnail, "AAMCAgADGQMAA9lnwP-e6mhUqlaqk5mT-Tg5mR8XNAACtmcAArpiCUrsexI73gMU6wEAB20AAzYE"},
	testbotdata.TestsInGroup2_bot:         {thumbnail, "AAMCAgADGQMAAxtnwP_gUCJBROB-DDn2hFazSffLjgACgGIAAgcVCErod_IyXuFt5wEAB20AAzYE"},
	testbotdata.TestsInGroup1_bot:         {thumbnail, "AAMCAgADGQMAAxRnwQABonSf1YDLopyPfyCqCmHvugIAAk5kAAKgtwlK9b3-oW_v7CABAAdtAAM2BA"},
	testbotdata.Make_event_bot:            {thumbnail, "AAMCAgADGQMAAxhnwQABxCuAaGFMzo7BBcXhIO-rOnIAAihzAAL9kglK51lwigpce_4BAAdtAAM2BA"},
	testbotdata.Testdatbot_bot:            {thumbnail, "AAMCAgADGQMAAgXyZ8EAAeQeQHVU5tPOoDGMqa_lPJn-AAJ-ZwACvNUJSrg3iGRYPgNcAQAHbQADNgQ"},
	testbotdata.SpeakOnlyWithAnton_bot:    {thumbnail, "AAMCAgADGQMAA1pnwQEJjrYPUETgD3qRu5CEPvXHkwACP3EAAnKmCUq6V6Jy2B47_QEAB20AAzYE"},
	testbotdata.LearnSpanishOrEnglish_bot: {thumbnail, "AAMCAgADGQMAAyhnwQEyYOJK--KcZ2lDlENDQtPWhwACHXMAAgRuCEoQS1g51A4uAgEAB20AAzYE"},
	testbotdata.Testmy_bots_bot:           {thumbnail, "AAMCAgADGQMAAxNnwQFJRUnvRLVPzhXNmbHGjuw2UAACNW4AAnrUCUr3_0vWo5KNSwEAB20AAzYE"},
}

var thumbvideo = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {thumbnail, "AAMCAgADGQMAAiSlZ8D5E5OMmfZKBbeefdv2Gri5IZkAAtJjAAJxXwhKRvX1CroOb6kBAAdtAAM2BA"},
	testbotdata.Tttesty_bot_bot:           {thumbnail, "AAMCAgADGQMAA9ZnwPlxEW-5pAmnY73Xoe0qkKRkcgACe3QAAnouCUqJ9p1gPKu7iwEAB20AAzYE"},
	testbotdata.TestsInGroup2_bot:         {thumbnail, "AAMCAgADGQMAAxlnwPmf-G47Rp96IsX301W6pKItcgACdmIAAgcVCErRV9i_faVuEQEAB20AAzYE"},
	testbotdata.TestsInGroup1_bot:         {thumbnail, "AAMCAgADGQMAAxFnwPnP4Y4P5dQvQz_1MB6_rGmOogACSWQAAqC3CUpMct8jS32-HAEAB20AAzYE"},
	testbotdata.Make_event_bot:            {thumbnail, "AAMCAgADGQMAAxRnwPnuQp_TO8CPlnjtzCE1yHonqwACFXMAAv2SCUp1GUK8PyBvJwEAB20AAzYE"},
	testbotdata.Testdatbot_bot:            {thumbnail, "AAMCAgADGQMAAgXtZ8D6DVpZ0hsLIrBF3_o4IyleVW4AAnJnAAK81QlK5TV-Fspt0oEBAAdtAAM2BA"},
	testbotdata.SpeakOnlyWithAnton_bot:    {thumbnail, "AAMCAgADGQMAA1dnwPpMrZ6g3llc5UkTSUqYK_453AACyHEAAkDFCErH8Ct30Sx8UAEAB20AAzYE"},
	testbotdata.LearnSpanishOrEnglish_bot: {thumbnail, "AAMCAgADGQMAAyVnwPp9EmHkqDYwhD9p7ImG-amf1wACg20AAtoBCEpQSVgEjpIUwgEAB20AAzYE"},
	testbotdata.Testmy_bots_bot:           {thumbnail, "AAMCAgADGQMAAxFnwPqYXVDgh4jQpOOqAZaGzwI_RgACMW4AAnrUCUolo-2dWkhSfgEAB20AAzYE"},
}

var thumbanim = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {thumbnail, "AAMCAgADGQMAAiSzZ8ECbRo8kzFett5P6Vu8Rg7KW_YAAt1jAAJxXwhKOiujh-pJhacBAAdtAAM2BA"},
	testbotdata.Tttesty_bot_bot:           {thumbnail, "AAMCAgADGQMAA9xnwQK4nKX-MdsuB9WP5j38Ixm73QACh3QAAnouCUqKaVnkNh4DLwEAB20AAzYE"},
	testbotdata.TestsInGroup2_bot:         {thumbnail, "AAMCAgADGQMAAx1nwQM1grCKRuwAAf-3d41AZ1jUUZgAAoViAAIHFQhKv_8UD7xtIEABAAdtAAM2BA"},
	testbotdata.TestsInGroup1_bot:         {thumbnail, "AAMCAgADGQMAAxZnwQNgNs1-aUDA3Gb5KKRPMFia0QACVWQAAqC3CUrv81zBuhzr4QEAB20AAzYE"},
	testbotdata.Make_event_bot:            {thumbnail, "AAMCAgADGQMAAxlnwQONVheRN51i4jzf4qNP0kfO_QACLXMAAv2SCUpWhdgkRylc6QEAB20AAzYE"},
	testbotdata.Testdatbot_bot:            {thumbnail, "AAMCAgADGQMAAgXzZ8EDs6XezloRAAGtjzD3zlgZFtRhAAKIZwACvNUJSuEnXo6P9XHMAQAHbQADNgQ"},
	testbotdata.SpeakOnlyWithAnton_bot:    {thumbnail, "AAMCAgADGQMAA1tnwQP3_8GMeBn48c0eKqSVHmtmzAACQHEAAnKmCUoC2kb5VDpw1AEAB20AAzYE"},
	testbotdata.LearnSpanishOrEnglish_bot: {thumbnail, "AAMCAgADGQMAAylnwQRIe3GsQN0CdVHD7njOp7YDCQACmm0AAtoBCEo6yq0Uu9X6pwEAB20AAzYE"},
	testbotdata.Testmy_bots_bot:           {thumbnail, "AAMCAgADGQMAAxRnwQSLB815QcTf0GmbGv8mjq-HgQACOW4AAnrUCUrRMhUSdn3wIgEAB20AAzYE"},
}

var thumbvideoN = map[string][]string{
	testbotdata.Secondtestbotforb_bot:     {thumbnail, "AAMCAgADGQMAAiS8Z8GS8tI_2EVTc0hzu-SWmg97uu8AAodoAAJxXwhKSJQfCFWzJ8IBAAdtAAM2BA"},
	testbotdata.Tttesty_bot_bot:           {thumbnail, "AAMCAgADGQMAA-BnwZONn1zPQi09IV6PUN1LJlWZawACUGYAAnouEUrRBalRsqar3QEAB20AAzYE"},
	testbotdata.TestsInGroup2_bot:         {thumbnail, "AAMCAgADGQMAAyFnwZOn2jlw_Y3IP464xz5ndmTnGgAC5GYAAgcVCEp7uzNqDtGGiAEAB20AAzYE"},
	testbotdata.TestsInGroup1_bot:         {thumbnail, "AAMCAgADGQMAAxxnwZP06Ietp0aIK_e0s_r0xzBNSQACSGkAAqC3CUqLC0MP-tfJYgEAB20AAzYE"},
	testbotdata.Make_event_bot:            {thumbnail, "AAMCAgADGQMAAyBnwZQYzDEKTXoKc-iB3-eqhAIivwACOWAAAv2SEUoT6oUWed2JYQEAB20AAzYE"},
	testbotdata.Testdatbot_bot:            {thumbnail, "AAMCAgADGQMAAgX2Z8GUWU9WqV-db6h0kUIbbeedJdYAAuhmAAK81RFKhVx1VVsoWZoBAAdtAAM2BA"},
	testbotdata.SpeakOnlyWithAnton_bot:    {thumbnail, "AAMCAgADGQMAA2BnwZRy4raBoVJ07Q9GAQioqCufqwACQV8AAnKmEUqUzrsxnAkT7gEAB20AAzYE"},
	testbotdata.LearnSpanishOrEnglish_bot: {thumbnail, "AAMCAgADGQMAAyxnwZSawrthDlNVnQz0lh_IHr40VQACD2wAAtoBEEoTKFHef3sgRQEAB20AAzYE"},
	testbotdata.Testmy_bots_bot:           {thumbnail, "AAMCAgADGQMAAxhnwZTWtQMb-MDS4KIkHm_ta589ywACwFsAAnrUEUpqkzYogveNbwEAB20AAzYE"},
}

var (
	parsemode       = []string{types.HTML, types.Markdown, types.MarkdownV2}
	groupnames      = []string{"A new group title/name", "What a cool name!", "something very very bad", "I don't like it at all"}
	titles          = []string{"the new admin", "something like a cool title", "don't have any clue", "to find a really cool title"}
	topicnames      = []string{"The Name", "No Name", "Something Like a Name", "Good Name", "Not a Good Name"}
	entities        = []*types.MessageEntity{{Offset: 1, Length: 4, Type: "italic"}}
	linkpopt        = &types.LinkPreviewOptions{IsDisabled: false, URL: "https://youtube.com"}
	chatperm        = &types.ChatPermissions{CanSendMessages: &f}
	botcomm         = []*types.BotCommand{{Command: "/start", Description: "Hello!"}, {Command: "/help", Description: "Help me!"}}
	botcommscope    = &types.BotCommandScope{Type: "default"}
	emojies         = []string{"😁", "😢", "😊", "😄", "😝", "😉", "🙈", "😜", "😞"}
	keywords        = []string{"Regular", "cool", "bad", "nothing", "etc"}
	pollOpt         = []*types.PollOption{{Text: "Yes!", VoterCount: 66}, {Text: "No!", VoterCount: 12}}
	kbnames         = []string{"inline", "reply-markup", "force-reply"}
	filenames       = []string{"storage", "telegram", "internet"}
	kb              = []func(*testcase){inlineKb, replyKb, forceKb}
	addToMapPhoto   = []func(*testcase){photoIntoMap, photoIntoMap, photoIntoMap}
	addToMapAudio   = []func(*testcase){audioIntoMap, audioIntoMap, audioIntoMap}
	addToMapDoc     = []func(*testcase){docIntoMap, docIntoMap, docIntoMap}
	addToMapVideo   = []func(*testcase){vidIntoMap, vidIntoMap, vidIntoMap}
	addToMapAnim    = []func(*testcase){animIntoMap, animIntoMap, animIntoMap}
	addToMapVoice   = []func(*testcase){vcIntoMap, vcIntoMap, vcIntoMap1}
	addToMapVdn     = []func(*testcase){vdnIntoMap, vdnIntoMap1, vdnIntoMap1}
	addToMapVdn1    = []func(*testcase){vdnIntoMap1, vdnIntoMap1, vdnIntoMap1}
	addToMapSticker = []func(*testcase){stckrIntoMap, stckrIntoMap, stckrIntoMap}
)

type mediagroup struct {
	amout     int
	photos    []formatter.IPhoto
	videos    []formatter.IVideo
	audios    []formatter.IAudio
	documents []formatter.IDocument
	phFunc    func(formatter.IPhoto) []func(string) error
	vdFunc    func(formatter.IVideo) []func(string) error
	adFunc    func(formatter.IAudio) []func(string) error
	docFunc   func(formatter.IDocument) []func(string) error
	all       bool
}

type testcase struct {
	msg           *formatter.Message
	prm           formatter.IParameters
	ch            formatter.IChat
	ph            formatter.IPhoto
	vd            formatter.IVideo
	an            formatter.IAnimation
	get           formatter.IGet
	ad            formatter.IAudio
	dc            formatter.IDocument
	vc            formatter.IVoice
	vdn           formatter.IVideoNote
	loc           formatter.ILocation
	con           formatter.IContact
	poll          formatter.IPoll
	link          formatter.ILink
	st            formatter.ISticker
	fr            formatter.IForum
	bot           formatter.IBot
	mg            *mediagroup
	addMedia      func(*testcase, *testing.T)
	mediaF        func(*testcase) []func(string) error
	mediadata     []string
	thumbdata     []string
	common        func(*testcase, *testcase, *testing.T)
	kbF           func(*testcase)
	thumb         bool
	thumbnailF    func(*testcase) []func(string) error
	code          [3]int
	errmsg        [3]string
	whattocheck   map[int]struct{}
	withoutString bool
	paid          bool
	token         string
	timetochange  chan struct{}
	workdone      chan struct{}
	getFuncs      []interface{}
}

func (tc *testcase) getReplyPrm() *types.ReplyParameters {
	return &types.ReplyParameters{MessageID: testbotdata.MessageID[tc.token], ChatID: chatid}
}

func (tc *testcase) changeToken(d, dd map[string][]string, t *testing.T) {
	for {
		select {
		case <-tc.workdone:
			return
		default:
			<-tc.timetochange
			tc.token = testbotdata.Tokens[rand.Intn(8)]
			if d != nil {
				tc.mediadata = d[tc.token]
			}
			if dd != nil {
				tc.thumbdata = dd[tc.token]
			}
		}

	}
}

func (tc *testcase) init() {
	// tc.token = testbotdata.Secondtestbotforb_bot
	// tc.token = testbotdata.Tttesty_bot_bot
	// tc.token = testbotdata.TestsInGroup2_bot
	// tc.token = testbotdata.TestsInGroup1_bot
	// tc.token = testbotdata.Make_event_bot
	// tc.token = testbotdata.Testdatbot_bot
	// tc.token = testbotdata.SpeakOnlyWithAnton_bot
	// tc.token = testbotdata.LearnSpanishOrEnglish_bot
	// tc.token = testbotdata.Testmy_bots_bot
	tc.token = testbotdata.Tokens[rand.Intn(8)]
	tc.msg = formatter.CreateEmpltyMessage()
	tc.prm = tc.msg.NewParameters()
	tc.ch = tc.msg.NewChat()
	tc.ph = tc.msg.NewPhoto()
	tc.ad = tc.msg.NewAudio()
	tc.dc = tc.msg.NewDocument()
	tc.vd = tc.msg.NewVideo()
	tc.an = tc.msg.NewAnimation()
	tc.vc = tc.msg.NewVoice()
	tc.vdn = tc.msg.NewVideoNote()
	tc.get = tc.msg.NewIGet()
	tc.loc = tc.msg.NewLocation()
	tc.con = tc.msg.NewContact()
	tc.poll = tc.msg.NewPoll()
	tc.link = tc.msg.NewLink()
	tc.st = tc.msg.NewSticker()
	tc.fr = tc.msg.NewForum()
	tc.bot = tc.msg.NewBot()
	tc.whattocheck = make(map[int]struct{})
	tc.timetochange = make(chan struct{})
	tc.workdone = make(chan struct{}, 1)
	tc.fillAllGetFunc()
	// types.DEBUG = true
	// logs.DebugOrInfo()
}

func (tc *testcase) fillAllGetFunc() {
	tc.getFuncs = make([]interface{}, 41)
	tc.getFuncs[0] = tc.get.Status
	tc.getFuncs[1] = tc.get.Error
	tc.getFuncs[2] = tc.get.Chat
	tc.getFuncs[3] = tc.get.Sender
	tc.getFuncs[4] = tc.get.Date
	tc.getFuncs[5] = tc.get.MessageID
	tc.getFuncs[6] = tc.get.MessageIDs
	tc.getFuncs[7] = tc.get.Replyed
	tc.getFuncs[8] = tc.get.ForwardOrigin
	tc.getFuncs[9] = tc.get.Photo
	tc.getFuncs[10] = tc.get.Audio
	tc.getFuncs[11] = tc.get.Document
	tc.getFuncs[12] = tc.get.Video
	tc.getFuncs[13] = tc.get.Animation
	tc.getFuncs[14] = tc.get.Voice
	tc.getFuncs[15] = tc.get.VideoNote
	tc.getFuncs[16] = tc.get.PaidMedia
	tc.getFuncs[17] = tc.get.MediaGroupID
	tc.getFuncs[18] = tc.get.Photos
	tc.getFuncs[19] = tc.get.Videos
	tc.getFuncs[20] = tc.get.Audios
	tc.getFuncs[21] = tc.get.Documents
	tc.getFuncs[22] = tc.get.Poll
	tc.getFuncs[23] = tc.get.Dice
	tc.getFuncs[24] = tc.get.ProfilePhotos
	tc.getFuncs[25] = tc.get.File
	tc.getFuncs[26] = tc.get.Stickers
	tc.getFuncs[27] = tc.get.Gifts
	tc.getFuncs[28] = tc.get.Message
	tc.getFuncs[29] = tc.get.String
	tc.getFuncs[30] = tc.get.InviteLink
	tc.getFuncs[31] = tc.get.ChatInfo
	tc.getFuncs[32] = tc.get.Members
	tc.getFuncs[33] = tc.get.Integer
	tc.getFuncs[34] = tc.get.Forum
	tc.getFuncs[35] = tc.get.Boosts
	tc.getFuncs[36] = tc.get.BusinessConnection
	tc.getFuncs[37] = tc.get.Commands
	tc.getFuncs[38] = tc.get.MenuButton
	tc.getFuncs[39] = tc.get.AdminRights
	tc.getFuncs[40] = tc.get.Messages
}

func (tc *testcase) checkResponse(j int) {
	for i := range tc.whattocheck {
		switch g := tc.getFuncs[i].(type) {
		case func() bool:
			if !g() {
				panic("Received boolean result has unexpectedly been 'false'. Should've been 'true'")
			}
		case func() (int, string):
			if code, msg := g(); (code != tc.code[i]) && (msg != tc.errmsg[i]) {
				panic(fmt.Sprintf("Received (int, string) result has unexpectedly been %d and %s. Should've been %d and `%s`", code, msg, tc.code[j], tc.errmsg[j]))
			}
		case func() string:
			if g() == "" {
				panic("Received string result has unexpectedly been empty. Should've been anything else but empty")
			}
		case func() int:
			if g() == 0 {
				panic("Received int result has unexpectedly been 0. Should've been anything else but 0")
			}
		case func() *int:
			if g() == nil {
				panic("Received *int result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []int:
			if g() == nil {
				panic("Received []int result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.Chat:
			if g() == nil {
				panic("Received *types.Chat result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.User:
			if g() == nil {
				panic("Received *types.User result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() formatter.IGet:
			if g() == nil {
				panic("Received formatter.IGet result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.MessageOrigin:
			if g() == nil {
				panic("Received *types.MessageOrigin result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.PhotoSize:
			if g() == nil {
				panic("Received []*types.PhotoSize result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.Audio:
			if g() == nil {
				panic("Received *types.Audio result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.Document:
			if g() == nil {
				panic("Received *types.Document result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.Video:
			if g() == nil {
				panic("Received *types.Video result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.Animation:
			if g() == nil {
				panic("Received *types.Animation result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.Voice:
			if g() == nil {
				panic("Received *types.Voice result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.VideoNote:
			if g() == nil {
				panic("Received *types.VideoNote result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.PaidMedia:
			if g() == nil {
				panic("Received *types.PaidMedia result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() [][]*types.PhotoSize:
			if g() == nil {
				panic("Received *[][]types.PhotoSize result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.Video:
			if g() == nil {
				panic("Received *[]types.Video result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.Audio:
			if g() == nil {
				panic("Received *[]types.Audio result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.Document:
			if g() == nil {
				panic("Received *[]types.Document result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.Poll:
			if g() == nil {
				panic("Received *types.Poll result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.Dice:
			if g() == nil {
				panic("Received *types.Dice result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.UserProfilePhotos:
			if g() == nil {
				panic("Received *types.UserProfilePhotos result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.File:
			if g() == nil {
				panic("Received *types.File result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.Sticker:
			if g() == nil {
				panic("Received *types.Sticker result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.Gift:
			if g() == nil {
				panic("Received *[]types.Gift result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.Message:
			if g() == nil {
				panic("Received *types.Message result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.Message:
			if g() == nil {
				panic("Received []*types.Message result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.ChatInviteLink:
			if g() == nil {
				panic("Received *types.ChatInviteLink result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.ChatFullInfo:
			if g() == nil {
				panic("Received *types.ChatFullInfo result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.ChatMember:
			if g() == nil {
				panic("Received *types.ChatMember result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.ForumTopic:
			if g() == nil {
				panic("Received *types.ForumTopic result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.ChatBoost:
			if g() == nil {
				panic("Received *types.ChatBoost result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.BusinessConnection:
			if g() == nil {
				panic("Received *types.BusinessConnection result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() []*types.BotCommand:
			if g() == nil {
				panic("Received *types.BotCommand result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.MenuButton:
			if g() == nil {
				panic("Received *types.MenuButton result has unexpectedly been nil. Should've been anything else but nil")
			}
		case func() *types.ChatAdministratorRights:
			if g() == nil {
				panic("Received *types.ChatAdministratorRights result has unexpectedly been nil. Should've been anything else but nil")
			}
		}
	}
}
