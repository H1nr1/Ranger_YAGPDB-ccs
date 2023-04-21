{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(inv(entory)?)(\s+|\z)`

    ©️ Ranger 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{/* Initiates variables */}}
{{$userID := .User.ID}}
{{$successColor := 0x00ff7b}}
{{$errorColor := 0xFF0000}}
{{$prefix := .ServerPrefix }}

{{/* Server shop */}}

{{/* Response */}}
{{$embed := sdict}}
{{$embed.Set "author" (sdict "name" $.User.Username "icon_url" ($.User.AvatarURL "1024"))}}
{{$embed.Set "timestamp" currentTime}}
{{with (dbGet 0 "EconomySettings")}}
    {{$a := sdict .Value}}
    {{$symbol := $a.symbol}}
	{{$userdata := or (dbGet $userID "userEconData").Value (sdict "inventory" sdict "streaks" (sdict "daily" 0 "weekly" 0 "monthly" 0))}}
	{{if $inventory := $userdata.inventory}} 
		{{$entry := cslice}}
		{{$field := cslice}}
		{{range $k,$v := $inventory}}
			{{$item := $k}}
			{{$desc := $v.desc}}
			{{$qty := $v.quantity}}
			{{$entry = $entry.Append (sdict "Name" $item "value" (joinStr "\n" (print "Description: " $desc) (print "Quantity: " (humanizeThousands $qty))) "inline" false)}}
		{{end}}
		{{$page := ""}}
		{{if $.CmdArgs}}
			{{$page = (index $.CmdArgs 0) | toInt}}
			{{if lt $page 1}}
				{{$page = 1}}
			{{end}}
		{{else}}
			{{$page = 1}}
		{{end}}
		{{$start := (mult 10 (sub $page 1))}}
		{{$stop := (mult $page 10)}}
		{{if ge $stop (len $entry)}}
			{{$stop = (len $entry)}}
		{{end}}
		{{if and (le $start $stop) (ge (len $entry) $start) (le $stop (len $entry))}}
			{{range (seq $start $stop)}}
				{{$field = $field.Append (index $entry .)}}
			{{end}}
		{{else}}
			{{$embed.Set "description" (print "This page is empty")}}
		{{end}}
        {{$embed.Set "title" (print "Inventory")}}
		{{$embed.Set "fields" $field}}
		{{$embed.Set "color" $successColor}}
		{{$embed.Set "footer" (sdict "text" (print "Page: " $page))}}
	{{else}}
		{{$embed.Set "description" (print "Your inventory is empty :(\nGet some items from the shop!")}}
		{{$embed.Set "color" $errorColor}}
	{{end}}
{{else}}
    {{$embed.Set "description" (print "No `Settings` database found.\nPlease set it up with the default values using `" $prefix "set default`")}}
    {{$embed.Set "color" $errorColor}}
{{end}}
{{sendMessage nil (cembed $embed)}}