{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Work`
©️ Ranger 2021
MIT License
*/}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$a := ""}}
{{$cash := ""}}
{{$b := .User.ID}}
{{if not (dbGet $b "EconomyInfo")}}
    {{dbSet .User.ID "EconomyInfo" (sdict "cash" 0 "bank" 0)}}
{{end}}
{{with (dbGet $b "EconomyInfo")}}
    {{$a = sdict .Value}}
    {{$cash = $a.cash}}
    {{$pay := (randInt 1 5000)}}
    {{$newCash := (add (toInt $cash) $pay)}}
    {{$workEmbed := (cembed
            "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))
            "description" (print "You decided to work today! You got paid a hefty " $pay)
            "color" 0x00ff7b
            "timestamp" currentTime
            )}}
    {{sendMessage nil $workEmbed}}
    {{$sdict := (dbGet .User.ID "EconomyInfo").Value}}
    {{$sdict.Set "cash" $newCash}}
    {{dbSet $b "EconomyInfo" $sdict}}
{{end}}