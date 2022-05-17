package blockToHtml

func findMarks(block Block, marks []string) (markDefs []MarkDef) {
	for _, mark := range marks {
		if mark == "strong" || mark == "em" {
			markDefs = append(markDefs, MarkDef{
				Key:  mark,
				Type: mark,
			})
		} else {
			for _, value := range block.MarkDefs {
				if value.Key == mark {
					markDefs = append(markDefs, value)
				}
			}
		}
	}
	return
}

func reverseMarks(input []MarkDef) []MarkDef {
	if len(input) == 0 {
		return input
	}
	return append(reverseMarks(input[1:]), input[0])
}

func parseChild(block Block) (html string) {
	for _, child := range block.Children {
		markDefs := findMarks(block, child.Marks)

		if len(markDefs) > 0 {
			for _, mark := range markDefs {
				if mark.Type == "strong" {
					html += `<strong class="block_content__strong">`
				} else if mark.Type == "em" {
					html += `<em class="block_content__emphasis">`
				} else if mark.Type == "link" {
					html += `<a class="block_content__link" href="` + mark.Href + `" target="_blank">`
				}
			}
			html += child.Text
			for _, mark := range reverseMarks(markDefs) {
				if mark.Type == "strong" {
					html += `</strong>`
				} else if mark.Type == "em" {
					html += `</em>`
				} else if mark.Type == "link" {
					html += `</a>`
				}
			}
		} else {
			html += child.Text
		}
	}

	return
}

func ToHtml(blocks []Block) (html string) {
	currentLevel, inList := 0, false

	for index, block := range blocks {
		if block.ListItem == "bullet" {
			if currentLevel < block.Level {
				currentLevel = block.Level
				inList = true
				html += `<ul class="block_content__list block_content__list--bullet">`
			} else if currentLevel > block.Level {
				currentLevel = block.Level
				inList = false
				html += "</ul>"
			}
		}

		if block.ListItem == "bullet" {
			html += `<li class="block_content__list_item">`
		}

		if block.Style == "normal" {
			html += `<p class="block_content__text">`
		} else if block.Style == "h1" {
			html += `<h1 class="block_content__title block_content__title--is-1">`
		} else if block.Style == "h2" {
			html += `<h2 class="block_content__title block_content__title--is-2">`
		} else if block.Style == "h3" {
			html += `<h3 class="block_content__title block_content__title--is-3">`
		} else if block.Style == "h4" {
			html += `<h4 class="block_content__title block_content__title--is-4">`
		} else if block.Style == "blockquote" {
			html += `<blockquote class="block_content__quote">`
		}

		html += parseChild(block)

		if block.Style == "normal" {
			html += "</p>"
		} else if block.Style == "h1" {
			html += "</h1>"
		} else if block.Style == "h2" {
			html += "</h2>"
		} else if block.Style == "h3" {
			html += "</h3>"
		} else if block.Style == "h4" {
			html += "</h4>"
		} else if block.Style == "blockquote" {
			html += "</blockquote>"
		}

		if block.ListItem == "bullet" {
			html += `</li>`
		}

		if inList && (blocks[len(blocks)-1].Key == block.Key || blocks[index+1].Level < currentLevel) {
			currentLevel = 0
			html += "</ul>"
		}
	}

	return
}
