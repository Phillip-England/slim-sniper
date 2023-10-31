package xlayout

import "fmt"

func Banner(title string, subtext string, menu string) string {
	overlay := Overlay()
	return fmt.Sprintf(`
		<div class='flex flex-row border-b items-center justify-between h-24 fixed top-0 z-40 w-full bg-white'>
			<div class='flex flex-col p-6'>
				<h1 class='text-2xl font-serif text-blue-700 font-semibold'>%s</h1>
				<p class='text-sm'>%s</p>
			</div>
			<div _='on click toggle .hidden on .top-nav-group' class='top-nav-group p-6'>
				<i class='fa-solid fa-bars fa-lg'></i>
			</div>
			<div _='on click toggle .hidden on .top-nav-group' class='top-nav-group p-6 hidden'>
				<i class='fa-solid fa-x fa-lg'></i>
			</div>
		</div>
		<div _='on click toggle .hidden on .top-nav-group' class='top-nav-group hidden'>%s</div>
		<div class='h-24'></div>
		<div class='top-nav-group hidden flex text-sm'>%s</div>
	`, title, subtext, overlay, menu)
}