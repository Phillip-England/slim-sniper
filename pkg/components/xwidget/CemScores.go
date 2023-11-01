package xwidget

import (
	"fmt"
	"net/url"
)

func CemScores(timescale string, store string) string {
	encodedTimescale := url.QueryEscape(timescale)
	store = url.QueryEscape(store)
	osatGet := fmt.Sprintf("/api/cem?timescale=%s&store=%s&metric=osat", encodedTimescale, store)
	tasteGet := fmt.Sprintf("/api/cem?timescale=%s&store=%s&metric=taste", encodedTimescale, store)
	aceGet := fmt.Sprintf("/api/cem?timescale=%s&store=%s&metric=ace", encodedTimescale, store)
	speedGet := fmt.Sprintf("/api/cem?timescale=%s&store=%s&metric=speed", encodedTimescale, store)
	cleanlinessGet := fmt.Sprintf("/api/cem?timescale=%s&store=%s&metric=cleanliness", encodedTimescale, store)
	accuracyGet := fmt.Sprintf("/api/cem?timescale=%s&store=%s&metric=accuracy", encodedTimescale, store)
	var southroadsOption string
	var uticaOption string
	var currentMonthOption string
	var threeMonthRollingOption string
	var yearToDateOption string
	if store == "southroads" {
		southroadsOption = "<option selected>southroads</option>"
	} else {
		southroadsOption = "<option>southroads</option>"
	}
	if store == "utica" {
		uticaOption = "<option selected>utica</option>"
	} else {
		uticaOption = "<option>utica</option>"
	}
	if timescale == "current month" {
		currentMonthOption = "<option selected>current month</option>"
	} else {
		currentMonthOption = "<option>current month</option>"
	}
	if timescale == "three month rolling" {
		threeMonthRollingOption = "<option selected>three month rolling</option>"
	} else {
		threeMonthRollingOption = "<option>three month rolling</option>"
	}
	if timescale == "year to date" {
		yearToDateOption = "<option selected>year to date</option>"
	} else {
		yearToDateOption = "<option>year to date</option>"
	}

	return fmt.Sprintf(`
	<form action='/team/cem' method='POST' class='bg-white rounded m-2 flex flex-col gap-6 p-6'>
			<div class='text-xs items-center justify-between gap-2 flex flex-row pb-2'>
				<h2 class='self-start text-lg'>CEM Results</h2>
				<div class='flex flex-col gap-2 text-xs justify-end items-end'>
					<select _='on click remove .hidden from .cem-result-search-button' name='store' class='border p-1 w-full'>
						%s%s
					</select>
					<select _='on click remove .hidden from .cem-result-search-button' name='timescale' class='border p-1'>
						%s%s%s
					</select>
				</div>
			</div>
			<ul class='flex flex-col rounded gap-2 text-xs'>
				<li class='flex flex-row justify-between border-b pb-4'>
					<div class='flex flex-row justify-between items-center gap-4'>
						<p>osat</p>	
					</div>
					<div hx-get=%s hx-swap='outerHTML' hx-trigger='load'></div>
				</li>
				<li class='flex flex-row justify-between border-b pb-4 pt-2'>
					<div class='flex flex-row items-center gap-4'>
						<p>taste</p>	
					</div>
					<div hx-get=%s hx-swap='outerHTML' hx-trigger='load'></div>
				</li>
				<li class='flex flex-row justify-between border-b pb-4 pt-2'>
					<div class='flex flex-row items-center gap-4'>
						<p>ace</p>	
					</div>
					<div hx-get=%s hx-swap='outerHTML' hx-trigger='load'></div>
				</li>
				<li class='flex flex-row justify-between border-b pb-4 pt-2'>
					<div class='flex flex-row items-center gap-4'>
						<p>speed</p>	
					</div>
					<div hx-get=%s hx-swap='outerHTML' hx-trigger='load'></div>
				</li>
				<li class='flex flex-row justify-between border-b pb-4 pt-2'>
					<div class='flex flex-row items-center gap-4'>
						<p>cleanliness</p>	
					</div>
					<div hx-get=%s hx-swap='outerHTML' hx-trigger='load'></div>
				</li>
				<li class='flex flex-row justify-between border-b pb-4 pt-2'>
					<div class='flex flex-row items-center gap-4'>
						<p>accuracy</p>	
					</div>
					<div hx-get=%s hx-swap='outerHTML' hx-trigger='load'></div>
				</li>
			</ul>
			<input value='Search' type='submit' class='cem-result-search-button hidden text-xs px-6 py-2 bg-blue-700 text-white rounded' />
		</form>
	`,southroadsOption, uticaOption, currentMonthOption, threeMonthRollingOption, yearToDateOption, osatGet, tasteGet, aceGet, speedGet, cleanlinessGet, accuracyGet)
}