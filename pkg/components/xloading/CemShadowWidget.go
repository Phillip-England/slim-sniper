package xloading

func CemShadowWidget() string {
	return `
		<div hx-get='/components/CustomerServiceScoreWidget' hx-swap='outerHTML' hx-trigger='load' class='bg-white m-2 rounded p-6 flex flex-col gap-4'>
			<div class='bg-gray-200 w-4/5 h-6 mb-2 animate-pulse rounded'></div>
			<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-4/5 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-1/2 h-4 animate-pulse rounded self-end'></div>
			<div class='bg-gray-200 w-2/3 h-4 animate-pulse rounded self-end'></div>
		</div>
	`
}