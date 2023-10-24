package xform

import "fmt"

func UpdateCemForm(formErr string) string {
	osatInput := TextInput("OSAT", "osat")
	tasteInput := TextInput("Taste", "taste")
	aceInput := TextInput("ACE", "ace")
	speedInput := TextInput("Speed", "speed")
	cleanlinessInput := TextInput("Cleanliness", "cleanliness")
	accuracyInput := TextInput("Accuracy", "accuracy")
	return fmt.Sprintf(`
		<form method='POST' action='/action/cem/update' class='bg-white p-6 m-2 rounded flex flex-col'>
			<div class='justify-between flex flex-row mb-4'>
				<h2 class='font-semibold text-lg'>Update CEM</h2>
				<select name='timescale' class='text-xs p-1 border '>
					<option value='current month'>Current Month</option>
					<option value='three month rolling'>Three Month Rolling</option>
					<option value='year to date'>Year To Date</option>
				</select>
			</div>
			<p class='text-xs mb-6 text-red-500'>%s</p>
			<div class='flex gap-2 flex-col mb-6'>%s%s%s%s%s%s</div>
			<input type='submit' class='text-xs bg-black p-2 w-1/3 rounded text-white' />
		</form>
	`, formErr, osatInput, tasteInput, aceInput,speedInput, cleanlinessInput, accuracyInput)
}