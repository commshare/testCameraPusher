package utils

import (
	"bytes"
	"golang.org/x/net/html/atom"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

func ConvertGBK(text1 string) string {
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(text1)), simplifiedchinese.GBK.NewDecoder()))
	text := string(data)
	return text
}

func ConvertUTF8(text1 string) string {
	data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(text1)), simplifiedchinese.GBK.NewEncoder()))
	text := string(data)
	return text
}

func GetAtomByString(a string) atom.Atom {

	switch a {
	case "A":
		return atom.A
	case "Abbr":
		return atom.Abbr
	case "Accept":
		return atom.Accept
	case "AcceptCharset":
		return atom.AcceptCharset
	case "Accesskey":
		return atom.Accesskey
	case "Action":
		return atom.Action
	case "Address":
		return atom.Address
	case "Align":
		return atom.Align
	case "Alt":
		return atom.Alt
	case "Annotation":
		return atom.Annotation
	case "AnnotationXml":
		return atom.AnnotationXml
	case "Applet":
		return atom.Applet
	case "Area":
		return atom.Area
	case "Article":
		return atom.Article
	case "Aside":
		return atom.Aside
	case "Async":
		return atom.Async
	case "Audio":
		return atom.Audio
	case "Autocomplete":
		return atom.Autocomplete

	}
	return 0

	/*
		Autofocus           Atom = 0xb309
		Autoplay            Atom = 0xce08
		B                   Atom = 0x101
		Base                Atom = 0xd604
		Basefont            Atom = 0xd608
		Bdi                 Atom = 0x1a03
		Bdo                 Atom = 0xe703
		Bgsound             Atom = 0x11807
		Big                 Atom = 0x12403
		Blink               Atom = 0x12705
		Blockquote          Atom = 0x12c0a
		Body                Atom = 0x2f04
		Br                  Atom = 0x202
		Button              Atom = 0x13606
		Canvas              Atom = 0x7f06
		Caption             Atom = 0x1bb07
		Center              Atom = 0x5b506
		Challenge           Atom = 0x21f09
		Charset             Atom = 0x2807
		Checked             Atom = 0x32807
		Cite                Atom = 0x3c804
		Class               Atom = 0x4de05
		Code                Atom = 0x14904
		Col                 Atom = 0x15003
		Colgroup            Atom = 0x15008
		Color               Atom = 0x15d05
		Cols                Atom = 0x16204
		Colspan             Atom = 0x16207
		Command             Atom = 0x17507
		Content             Atom = 0x42307
		Contenteditable     Atom = 0x4230f
		Contextmenu         Atom = 0x3310b
		Controls            Atom = 0x18808
		Coords              Atom = 0x19406
		Crossorigin         Atom = 0x19f0b
		Data                Atom = 0x44a04
		Datalist            Atom = 0x44a08
		Datetime            Atom = 0x23c08
		Dd                  Atom = 0x26702
		Default             Atom = 0x8607
		Defer               Atom = 0x14b05
		Del                 Atom = 0x3ef03
		Desc                Atom = 0x4db04
		Details             Atom = 0x4807
		Dfn                 Atom = 0x6103
		Dialog              Atom = 0x1b06
		Dir                 Atom = 0x6903
		Dirname             Atom = 0x6907
		Disabled            Atom = 0x10c08
		Div                 Atom = 0x11303
		Dl                  Atom = 0x11e02
		Download            Atom = 0x40008
		Draggable           Atom = 0x17b09
		Dropzone            Atom = 0x39108
		Dt                  Atom = 0x50902
		Em                  Atom = 0x6502
		Embed               Atom = 0x6505
		Enctype             Atom = 0x21107
		Face                Atom = 0x5b304
		Fieldset            Atom = 0x1b008
		Figcaption          Atom = 0x1b80a
		Figure              Atom = 0x1cc06
		Font                Atom = 0xda04
		Footer              Atom = 0x8d06
		For                 Atom = 0x1d803
		ForeignObject       Atom = 0x1d80d
		Foreignobject       Atom = 0x1e50d
		Form                Atom = 0x1f204
		Formaction          Atom = 0x1f20a
		Formenctype         Atom = 0x20d0b
		Formmethod          Atom = 0x2280a
		Formnovalidate      Atom = 0x2320e
		Formtarget          Atom = 0x2470a
		Frame               Atom = 0x9a05
		Frameset            Atom = 0x9a08
		H1                  Atom = 0x26e02
		H2                  Atom = 0x29402
		H3                  Atom = 0x2a702
		H4                  Atom = 0x2e902
		H5                  Atom = 0x2f302
		H6                  Atom = 0x50b02
		Head                Atom = 0x2d504
		Header              Atom = 0x2d506
		Headers             Atom = 0x2d507
		Height              Atom = 0x25106
		Hgroup              Atom = 0x25906
		Hidden              Atom = 0x26506
		High                Atom = 0x26b04
		Hr                  Atom = 0x27002
		Href                Atom = 0x27004
		Hreflang            Atom = 0x27008
		Html                Atom = 0x25504
		HttpEquiv           Atom = 0x2780a
		I                   Atom = 0x601
		Icon                Atom = 0x42204
		Id                  Atom = 0x8502
		Iframe              Atom = 0x29606
		Image               Atom = 0x29c05
		Img                 Atom = 0x2a103
		Input               Atom = 0x3e805
		Inputmode           Atom = 0x3e809
		Ins                 Atom = 0x1a803
		Isindex             Atom = 0x2a907
		Ismap               Atom = 0x2b005
		Itemid              Atom = 0x33c06
		Itemprop            Atom = 0x3c908
		Itemref             Atom = 0x5ad07
		Itemscope           Atom = 0x2b909
		Itemtype            Atom = 0x2c308
		Kbd                 Atom = 0x1903
		Keygen              Atom = 0x3906
		Keytype             Atom = 0x53707
		Kind                Atom = 0x10904
		Label               Atom = 0xf005
		Lang                Atom = 0x27404
		Legend              Atom = 0x18206
		Li                  Atom = 0x1202
		Link                Atom = 0x12804
		List                Atom = 0x44e04
		Listing             Atom = 0x44e07
		Loop                Atom = 0xf404
		Low                 Atom = 0x11f03
		Malignmark          Atom = 0x100a
		Manifest            Atom = 0x5f108
		Map                 Atom = 0x2b203
		Mark                Atom = 0x1604
		Marquee             Atom = 0x2cb07
		Math                Atom = 0x2d204
		Max                 Atom = 0x2e103
		Maxlength           Atom = 0x2e109
		Media               Atom = 0x6e05
		Mediagroup          Atom = 0x6e0a
		Menu                Atom = 0x33804
		Menuitem            Atom = 0x33808
		Meta                Atom = 0x45d04
		Meter               Atom = 0x24205
		Method              Atom = 0x22c06
		Mglyph              Atom = 0x2a206
		Mi                  Atom = 0x2eb02
		Min                 Atom = 0x2eb03
		Minlength           Atom = 0x2eb09
		Mn                  Atom = 0x23502
		Mo                  Atom = 0x3ed02
		Ms                  Atom = 0x2bc02
		Mtext               Atom = 0x2f505
		Multiple            Atom = 0x30308
		Muted               Atom = 0x30b05
		Name                Atom = 0x6c04
		Nav                 Atom = 0x3e03
		Nobr                Atom = 0x5704
		Noembed             Atom = 0x6307
		Noframes            Atom = 0x9808
		Noscript            Atom = 0x3d208
		Novalidate          Atom = 0x2360a
		Object              Atom = 0x1ec06
		Ol                  Atom = 0xc902
		Onabort             Atom = 0x13a07
		Onafterprint        Atom = 0x1c00c
		Onautocomplete      Atom = 0x1fa0e
		Onautocompleteerror Atom = 0x1fa13
		Onbeforeprint       Atom = 0x6040d
		Onbeforeunload      Atom = 0x4e70e
		Onblur              Atom = 0xaa06
		Oncancel            Atom = 0xe908
		Oncanplay           Atom = 0x28509
		Oncanplaythrough    Atom = 0x28510
		Onchange            Atom = 0x3a708
		Onclick             Atom = 0x31007
		Onclose             Atom = 0x31707
		Oncontextmenu       Atom = 0x32f0d
		Oncuechange         Atom = 0x3420b
		Ondblclick          Atom = 0x34d0a
		Ondrag              Atom = 0x35706
		Ondragend           Atom = 0x35709
		Ondragenter         Atom = 0x3600b
		Ondragleave         Atom = 0x36b0b
		Ondragover          Atom = 0x3760a
		Ondragstart         Atom = 0x3800b
		Ondrop              Atom = 0x38f06
		Ondurationchange    Atom = 0x39f10
		Onemptied           Atom = 0x39609
		Onended             Atom = 0x3af07
		Onerror             Atom = 0x3b607
		Onfocus             Atom = 0x3bd07
		Onhashchange        Atom = 0x3da0c
		Oninput             Atom = 0x3e607
		Oninvalid           Atom = 0x3f209
		Onkeydown           Atom = 0x3fb09
		Onkeypress          Atom = 0x4080a
		Onkeyup             Atom = 0x41807
		Onlanguagechange    Atom = 0x43210
		Onload              Atom = 0x44206
		Onloadeddata        Atom = 0x4420c
		Onloadedmetadata    Atom = 0x45510
		Onloadstart         Atom = 0x46b0b
		Onmessage           Atom = 0x47609
		Onmousedown         Atom = 0x47f0b
		Onmousemove         Atom = 0x48a0b
		Onmouseout          Atom = 0x4950a
		Onmouseover         Atom = 0x4a20b
		Onmouseup           Atom = 0x4ad09
		Onmousewheel        Atom = 0x4b60c
		Onoffline           Atom = 0x4c209
		Ononline            Atom = 0x4cb08
		Onpagehide          Atom = 0x4d30a
		Onpageshow          Atom = 0x4fe0a
		Onpause             Atom = 0x50d07
		Onplay              Atom = 0x51706
		Onplaying           Atom = 0x51709
		Onpopstate          Atom = 0x5200a
		Onprogress          Atom = 0x52a0a
		Onratechange        Atom = 0x53e0c
		Onreset             Atom = 0x54a07
		Onresize            Atom = 0x55108
		Onscroll            Atom = 0x55f08
		Onseeked            Atom = 0x56708
		Onseeking           Atom = 0x56f09
		Onselect            Atom = 0x57808
		Onshow              Atom = 0x58206
		Onsort              Atom = 0x58b06
		Onstalled           Atom = 0x59509
		Onstorage           Atom = 0x59e09
		Onsubmit            Atom = 0x5a708
		Onsuspend           Atom = 0x5bb09
		Ontimeupdate        Atom = 0xdb0c
		Ontoggle            Atom = 0x5c408
		Onunload            Atom = 0x5cc08
		Onvolumechange      Atom = 0x5d40e
		Onwaiting           Atom = 0x5e209
		Open                Atom = 0x3cf04
		Optgroup            Atom = 0xf608
		Optimum             Atom = 0x5eb07
		Option              Atom = 0x60006
		Output              Atom = 0x49c06
		P                   Atom = 0xc01
		Param               Atom = 0xc05
		Pattern             Atom = 0x5107
		Ping                Atom = 0x7704
		Placeholder         Atom = 0xc30b
		Plaintext           Atom = 0xfd09
		Poster              Atom = 0x15706
		Pre                 Atom = 0x25e03
		Preload             Atom = 0x25e07
		Progress            Atom = 0x52c08
		Prompt              Atom = 0x5fa06
		Public              Atom = 0x41e06
		Q                   Atom = 0x13101
		Radiogroup          Atom = 0x30a
		Readonly            Atom = 0x2fb08
		Rel                 Atom = 0x25f03
		Required            Atom = 0x1d008
		Reversed            Atom = 0x5a08
		Rows                Atom = 0x9204
		Rowspan             Atom = 0x9207
		Rp                  Atom = 0x1c602
		Rt                  Atom = 0x13f02
		Ruby                Atom = 0xaf04
		S                   Atom = 0x2c01
		Samp                Atom = 0x4e04
		Sandbox             Atom = 0xbb07
		Scope               Atom = 0x2bd05
		Scoped              Atom = 0x2bd06
		Script              Atom = 0x3d406
		Seamless            Atom = 0x31c08
		Section             Atom = 0x4e207
		Select              Atom = 0x57a06
		Selected            Atom = 0x57a08
		Shape               Atom = 0x4f905
		Size                Atom = 0x55504
		Sizes               Atom = 0x55505
		Small               Atom = 0x18f05
		Sortable            Atom = 0x58d08
		Sorted              Atom = 0x19906
		Source              Atom = 0x1aa06
		Spacer              Atom = 0x2db06
		Span                Atom = 0x9504
		Spellcheck          Atom = 0x3230a
		Src                 Atom = 0x3c303
		Srcdoc              Atom = 0x3c306
		Srclang             Atom = 0x41107
		Start               Atom = 0x38605
		Step                Atom = 0x5f704
		Strike              Atom = 0x53306
		Strong              Atom = 0x55906
		Style               Atom = 0x61105
		Sub                 Atom = 0x5a903
		Summary             Atom = 0x61607
		Sup                 Atom = 0x61d03
		Svg                 Atom = 0x62003
		System              Atom = 0x62306
		Tabindex            Atom = 0x46308
		Table               Atom = 0x42d05
		Target              Atom = 0x24b06
		Tbody               Atom = 0x2e05
		Td                  Atom = 0x4702
		Template            Atom = 0x62608
		Textarea            Atom = 0x2f608
		Tfoot               Atom = 0x8c05
		Th                  Atom = 0x22e02
		Thead               Atom = 0x2d405
		Time                Atom = 0xdd04
		Title               Atom = 0xa105
		Tr                  Atom = 0x10502
		Track               Atom = 0x10505
		Translate           Atom = 0x14009
		Tt                  Atom = 0x5302
		Type                Atom = 0x21404
		Typemustmatch       Atom = 0x2140d
		U                   Atom = 0xb01
		Ul                  Atom = 0x8a02
		Usemap              Atom = 0x51106
		Value               Atom = 0x4005
		Var                 Atom = 0x11503
		Video               Atom = 0x28105
		Wbr                 Atom = 0x12103
		Width               Atom = 0x50705
		Wrap                Atom = 0x58704
		Xmp                 Atom = 0xc103*/
}
