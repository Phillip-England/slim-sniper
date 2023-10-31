package xform

import "fmt"

func UpdateCemForm(formErr string) string {
	return fmt.Sprintf(`
		<form method='POST' action='/action/cem/update' class='bg-white p-6 m-2 rounded flex flex-col'>
			<div class='justify-between flex flex-row mb-2'>
				<h2 class='text-lg'>Update CEM</h2>
				<div class='flex flex-col gap-2'>
					<select name='store' class='text-xs p-1 border '>
						<option value='southroads'>Southroads</option>
						<option value='utica'>Utica</option>
					</select>
					<select name='timescale' class='text-xs p-1 border '>
						<option value='current month'>Current Month</option>
						<option value='three month rolling'>Three Month Rolling</option>
						<option value='year to date'>Year To Date</option>
						<option value='year to date'>Goal</option>
					</select>
				</div>
			</div>
			<p class='text-xs mb-6 text-red-500'>%s</p>
			<div class='flex gap-2 flex-col mb-6'>
				<div class='flex flex-row gap-4'>
					<div class='flex flex-col text-xs gap-1'>
						<label for='osat'>osat</label>
						<input name='osat' type='text' class='border p-1' />
					</div>
					<div class='flex flex-col text-xs gap-1'>
						<label for='taste'>taste</label>
						<input name='taste' type='text' class='border p-1' />
					</div>
				</div>
				<div class='flex flex-row gap-4'>
					<div class='flex flex-col text-xs gap-1'>
						<label for='ace'>ace</label>
						<input name='ace' type='text' class='border p-1' />
					</div>
					<div class='flex flex-col text-xs gap-1'>
						<label for='speed'>speed</label>
						<input name='speed' type='text' class='border p-1' />
					</div>
				</div>
				<div class='flex flex-row gap-4'>
					<div class='flex flex-col text-xs gap-1'>
						<label for='cleanliness'>cleanliness</label>
						<input name='cleanliness' type='text' class='border p-1' />
					</div>
					<div class='flex flex-col text-xs gap-1'>
						<label for='accuracy'>accuracy</label>
						<input name='accuracy' type='text' class='border p-1' />
					</div>
				</div>
			</div>
			<input type='submit' class='mt-2 text-xs bg-blue-700 p-2 rounded text-white' />
		</form>
	`, formErr)
}