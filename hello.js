// based on https://www.masterduelmeta.com/deck-tester, handle exceptions for unreleased cards
copy(["main", "extra"].flatMap(
	section => [...(document.getElementById(section)?.children[1]?.children || [])].map(
		function (child) {
			v = child.children[0].children[0]
			if (v) {
				return v.title
			} else {
				return "Ancient Gear Statue" // unreleased card name
			}
		}))
)
