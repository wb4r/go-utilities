///////////////////
// 	GENERAL STUFF
///////////////////

func blank_line() {
	fmt.Println("")
}

//////////////////////////////
// VARIABLES AND ASSIGNMENTS
//////////////////////////////


/////////////////
// CONTROL FLOW
/////////////////


//////////////////////
// SLICES AND ARRAYS
//////////////////////

func arr_remove_int(given_slice []int, i int) []int {
	return append(given_slice[:i], given_slice[i+1:]...)
}
func arr_remove_str(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

///////////////
// 	  MAPS
///////////////


///////////////
//  FUNCTIONS
///////////////


////////////////
//   POINTERS
////////////////

































































































//
