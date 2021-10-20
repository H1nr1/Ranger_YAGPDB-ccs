{{/* 
        Case system by Maverick Wolf (549820835230253060)
        Made & modified by Ranger (765316548516380732)

    **Tools & Util > Moderation > Mute > Mute DM**
©️ Ranger 2021
MIT License
*/}}


{{/*        Notes
    Please be aware that even though the original custom commands had the time, I removed it from this for NOW. They may be added back if I can get it to show the current time, not the time in GMT.
*/}}

{{/* Configuration values start */}}
{{$LogChannel := 784132358085017604}} {{/* Modlog channel */}}
{{$dm := 1}} {{/* change to 0 if you don't wanna DM the offender about the moderation action */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$icon := ""}}
{{$name := printf "%s (%d)" .Guild.Name .Guild.ID}}
{{if .Guild.Icon}}
	{{$ext := "webp"}}
	{{if eq (slice .Guild.Icon 0 2) "a_"}}
        {{$ext = "gif"}}
    {{end}}
	{{$icon = printf "https://cdn.discordapp.com/icons/%d/%s.%s" .Guild.ID .Guild.Icon $ext}}
{{end}}

{{$id := .User.ID}}

{{$channel2 := $LogChannel}}
{{if .Channel.ID}}
    {{$channel2 = .Channel.ID}}
{{end}}

{{/* log & DM messages */}}
{{if $dm}}
    {{$UnmuteDM := cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
            "description" (print "**Server:** " .Guild.Name "\n**Action:** `Unmute`\n**Reason: **" (joinStr " " (split (reReplace `Automoderator:` .Reason "<:Bot:787563190221406259>:") "\n")))
            "thumbnail" (sdict "url" $icon)
            "color" 3553599
            "timestamp" currentTime
            }}
    {{sendDM $UnmuteDM}}
{{end}}

{{$x := sendMessageRetID $LogChannel (cembed
            "author" (sdict "icon_url" (.Author.AvatarURL "1024") "name" (print .Author.String " (ID " .Author.ID ")"))
            "description" (print "<:Management:788937280508657694> **Who:** " .User.Mention " `(ID " .User.ID ")`\n<:Metadata:788937280508657664> **Action:** `Unmute`\n<:Assetlibrary:788937280554926091> **Channel:** <#" $channel2 ">\n<:Manifest:788937280579698728> **Reason:** " (joinStr " " (split (reReplace `Automoderator:` .Reason "<:Bot:787563190221406259>:") "\n")))
            "thumbnail" (sdict "url" (.User.AvatarURL "256"))
            "color" 6473311
            )}}

{{$Response := sendMessageRetID nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print "Case type: Unmute"))
            "description" (print .Author.Mention " Has successfully unmuted " .User.Mention " :thumbsup:")
            "color" 3553599
            "timestamp" currentTime
            )}}
{{deleteMessage nil $Response 5}}