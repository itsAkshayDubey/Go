// pronicnumber_test.go
// description: Returns true if number is pronic, false otherwise
// author: Akshay Dubey (https://github.com/itsAkshayDubey)
// see pronicnumber.go

package math_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/math"
)

func TestPronicNumber(t *testing.T) {
	var tests = []struct {
		name          string
		n             int
		expectedValue bool
	}{
		{"n = -12", -12, false},
		{"n = 0", 0, true},
		{"n = 1", 1, false},
		{"n = 2", 2, true},
		{"n = 3", 3, false},
		{"n = 4", 4, false},
		{"n = 5", 5, false},
		{"n = 6", 6, true},
		{"n = 7", 7, false},
		{"n = 8", 8, false},
		{"n = 9", 9, false},
		{"n = 10", 10, false},
		{"n = 11", 11, false},
		{"n = 12", 12, true},
		{"n = 13", 13, false},
		{"n = 14", 14, false},
		{"n = 15", 15, false},
		{"n = 16", 16, false},
		{"n = 17", 17, false},
		{"n = 18", 18, false},
		{"n = 19", 19, false},
		{"n = 20", 20, true},
		{"n = 21", 21, false},
		{"n = 22", 22, false},
		{"n = 23", 23, false},
		{"n = 24", 24, false},
		{"n = 25", 25, false},
		{"n = 26", 26, false},
		{"n = 27", 27, false},
		{"n = 28", 28, false},
		{"n = 29", 29, false},
		{"n = 30", 30, true},
		{"n = 31", 31, false},
		{"n = 32", 32, false},
		{"n = 33", 33, false},
		{"n = 34", 34, false},
		{"n = 35", 35, false},
		{"n = 36", 36, false},
		{"n = 37", 37, false},
		{"n = 38", 38, false},
		{"n = 39", 39, false},
		{"n = 40", 40, false},
		{"n = 41", 41, false},
		{"n = 42", 42, true},
		{"n = 43", 43, false},
		{"n = 44", 44, false},
		{"n = 45", 45, false},
		{"n = 46", 46, false},
		{"n = 47", 47, false},
		{"n = 48", 48, false},
		{"n = 49", 49, false},
		{"n = 50", 50, false},
		{"n = 51", 51, false},
		{"n = 52", 52, false},
		{"n = 53", 53, false},
		{"n = 54", 54, false},
		{"n = 55", 55, false},
		{"n = 56", 56, true},
		{"n = 57", 57, false},
		{"n = 58", 58, false},
		{"n = 59", 59, false},
		{"n = 60", 60, false},
		{"n = 61", 61, false},
		{"n = 62", 62, false},
		{"n = 63", 63, false},
		{"n = 64", 64, false},
		{"n = 65", 65, false},
		{"n = 66", 66, false},
		{"n = 67", 67, false},
		{"n = 68", 68, false},
		{"n = 69", 69, false},
		{"n = 70", 70, false},
		{"n = 71", 71, false},
		{"n = 72", 72, true},
		{"n = 73", 73, false},
		{"n = 74", 74, false},
		{"n = 75", 75, false},
		{"n = 76", 76, false},
		{"n = 77", 77, false},
		{"n = 78", 78, false},
		{"n = 79", 79, false},
		{"n = 80", 80, false},
		{"n = 81", 81, false},
		{"n = 82", 82, false},
		{"n = 83", 83, false},
		{"n = 84", 84, false},
		{"n = 85", 85, false},
		{"n = 86", 86, false},
		{"n = 87", 87, false},
		{"n = 88", 88, false},
		{"n = 89", 89, false},
		{"n = 90", 90, true},
		{"n = 91", 91, false},
		{"n = 92", 92, false},
		{"n = 93", 93, false},
		{"n = 94", 94, false},
		{"n = 95", 95, false},
		{"n = 96", 96, false},
		{"n = 97", 97, false},
		{"n = 98", 98, false},
		{"n = 99", 99, false},
		{"n = 100", 100, false},
		{"n = 101", 101, false},
		{"n = 102", 102, false},
		{"n = 103", 103, false},
		{"n = 104", 104, false},
		{"n = 105", 105, false},
		{"n = 106", 106, false},
		{"n = 107", 107, false},
		{"n = 108", 108, false},
		{"n = 109", 109, false},
		{"n = 110", 110, true},
		{"n = 111", 111, false},
		{"n = 112", 112, false},
		{"n = 113", 113, false},
		{"n = 114", 114, false},
		{"n = 115", 115, false},
		{"n = 116", 116, false},
		{"n = 117", 117, false},
		{"n = 118", 118, false},
		{"n = 119", 119, false},
		{"n = 120", 120, false},
		{"n = 121", 121, false},
		{"n = 122", 122, false},
		{"n = 123", 123, false},
		{"n = 124", 124, false},
		{"n = 125", 125, false},
		{"n = 126", 126, false},
		{"n = 127", 127, false},
		{"n = 128", 128, false},
		{"n = 129", 129, false},
		{"n = 130", 130, false},
		{"n = 131", 131, false},
		{"n = 132", 132, true},
		{"n = 133", 133, false},
		{"n = 134", 134, false},
		{"n = 135", 135, false},
		{"n = 136", 136, false},
		{"n = 137", 137, false},
		{"n = 138", 138, false},
		{"n = 139", 139, false},
		{"n = 140", 140, false},
		{"n = 141", 141, false},
		{"n = 142", 142, false},
		{"n = 143", 143, false},
		{"n = 144", 144, false},
		{"n = 145", 145, false},
		{"n = 146", 146, false},
		{"n = 147", 147, false},
		{"n = 148", 148, false},
		{"n = 149", 149, false},
		{"n = 150", 150, false},
		{"n = 151", 151, false},
		{"n = 152", 152, false},
		{"n = 153", 153, false},
		{"n = 154", 154, false},
		{"n = 155", 155, false},
		{"n = 156", 156, true},
		{"n = 157", 157, false},
		{"n = 158", 158, false},
		{"n = 159", 159, false},
		{"n = 160", 160, false},
		{"n = 161", 161, false},
		{"n = 162", 162, false},
		{"n = 163", 163, false},
		{"n = 164", 164, false},
		{"n = 165", 165, false},
		{"n = 166", 166, false},
		{"n = 167", 167, false},
		{"n = 168", 168, false},
		{"n = 169", 169, false},
		{"n = 170", 170, false},
		{"n = 171", 171, false},
		{"n = 172", 172, false},
		{"n = 173", 173, false},
		{"n = 174", 174, false},
		{"n = 175", 175, false},
		{"n = 176", 176, false},
		{"n = 177", 177, false},
		{"n = 178", 178, false},
		{"n = 179", 179, false},
		{"n = 180", 180, false},
		{"n = 181", 181, false},
		{"n = 182", 182, true},
		{"n = 183", 183, false},
		{"n = 184", 184, false},
		{"n = 185", 185, false},
		{"n = 186", 186, false},
		{"n = 187", 187, false},
		{"n = 188", 188, false},
		{"n = 189", 189, false},
		{"n = 190", 190, false},
		{"n = 191", 191, false},
		{"n = 192", 192, false},
		{"n = 193", 193, false},
		{"n = 194", 194, false},
		{"n = 195", 195, false},
		{"n = 196", 196, false},
		{"n = 197", 197, false},
		{"n = 198", 198, false},
		{"n = 199", 199, false},
		{"n = 200", 200, false},
		{"n = 201", 201, false},
		{"n = 202", 202, false},
		{"n = 203", 203, false},
		{"n = 204", 204, false},
		{"n = 205", 205, false},
		{"n = 206", 206, false},
		{"n = 207", 207, false},
		{"n = 208", 208, false},
		{"n = 209", 209, false},
		{"n = 210", 210, true},
		{"n = 211", 211, false},
		{"n = 212", 212, false},
		{"n = 213", 213, false},
		{"n = 214", 214, false},
		{"n = 215", 215, false},
		{"n = 216", 216, false},
		{"n = 217", 217, false},
		{"n = 218", 218, false},
		{"n = 219", 219, false},
		{"n = 220", 220, false},
		{"n = 221", 221, false},
		{"n = 222", 222, false},
		{"n = 223", 223, false},
		{"n = 224", 224, false},
		{"n = 225", 225, false},
		{"n = 226", 226, false},
		{"n = 227", 227, false},
		{"n = 228", 228, false},
		{"n = 229", 229, false},
		{"n = 230", 230, false},
		{"n = 231", 231, false},
		{"n = 232", 232, false},
		{"n = 233", 233, false},
		{"n = 234", 234, false},
		{"n = 235", 235, false},
		{"n = 236", 236, false},
		{"n = 237", 237, false},
		{"n = 238", 238, false},
		{"n = 239", 239, false},
		{"n = 240", 240, true},
		{"n = 241", 241, false},
		{"n = 242", 242, false},
		{"n = 243", 243, false},
		{"n = 244", 244, false},
		{"n = 245", 245, false},
		{"n = 246", 246, false},
		{"n = 247", 247, false},
		{"n = 248", 248, false},
		{"n = 249", 249, false},
		{"n = 250", 250, false},
		{"n = 251", 251, false},
		{"n = 252", 252, false},
		{"n = 253", 253, false},
		{"n = 254", 254, false},
		{"n = 255", 255, false},
		{"n = 256", 256, false},
		{"n = 257", 257, false},
		{"n = 258", 258, false},
		{"n = 259", 259, false},
		{"n = 260", 260, false},
		{"n = 261", 261, false},
		{"n = 262", 262, false},
		{"n = 263", 263, false},
		{"n = 264", 264, false},
		{"n = 265", 265, false},
		{"n = 266", 266, false},
		{"n = 267", 267, false},
		{"n = 268", 268, false},
		{"n = 269", 269, false},
		{"n = 270", 270, false},
		{"n = 271", 271, false},
		{"n = 272", 272, true},
		{"n = 273", 273, false},
		{"n = 274", 274, false},
		{"n = 275", 275, false},
		{"n = 276", 276, false},
		{"n = 277", 277, false},
		{"n = 278", 278, false},
		{"n = 279", 279, false},
		{"n = 280", 280, false},
		{"n = 281", 281, false},
		{"n = 282", 282, false},
		{"n = 283", 283, false},
		{"n = 284", 284, false},
		{"n = 285", 285, false},
		{"n = 286", 286, false},
		{"n = 287", 287, false},
		{"n = 288", 288, false},
		{"n = 289", 289, false},
		{"n = 290", 290, false},
		{"n = 291", 291, false},
		{"n = 292", 292, false},
		{"n = 293", 293, false},
		{"n = 294", 294, false},
		{"n = 295", 295, false},
		{"n = 296", 296, false},
		{"n = 297", 297, false},
		{"n = 298", 298, false},
		{"n = 299", 299, false},
		{"n = 300", 300, false},
		{"n = 301", 301, false},
		{"n = 302", 302, false},
		{"n = 303", 303, false},
		{"n = 304", 304, false},
		{"n = 305", 305, false},
		{"n = 306", 306, true},
		{"n = 307", 307, false},
		{"n = 308", 308, false},
		{"n = 309", 309, false},
		{"n = 310", 310, false},
		{"n = 311", 311, false},
		{"n = 312", 312, false},
		{"n = 313", 313, false},
		{"n = 314", 314, false},
		{"n = 315", 315, false},
		{"n = 316", 316, false},
		{"n = 317", 317, false},
		{"n = 318", 318, false},
		{"n = 319", 319, false},
		{"n = 320", 320, false},
		{"n = 321", 321, false},
		{"n = 322", 322, false},
		{"n = 323", 323, false},
		{"n = 324", 324, false},
		{"n = 325", 325, false},
		{"n = 326", 326, false},
		{"n = 327", 327, false},
		{"n = 328", 328, false},
		{"n = 329", 329, false},
		{"n = 330", 330, false},
		{"n = 331", 331, false},
		{"n = 332", 332, false},
		{"n = 333", 333, false},
		{"n = 334", 334, false},
		{"n = 335", 335, false},
		{"n = 336", 336, false},
		{"n = 337", 337, false},
		{"n = 338", 338, false},
		{"n = 339", 339, false},
		{"n = 340", 340, false},
		{"n = 341", 341, false},
		{"n = 342", 342, true},
		{"n = 343", 343, false},
		{"n = 344", 344, false},
		{"n = 345", 345, false},
		{"n = 346", 346, false},
		{"n = 347", 347, false},
		{"n = 348", 348, false},
		{"n = 349", 349, false},
		{"n = 350", 350, false},
		{"n = 351", 351, false},
		{"n = 352", 352, false},
		{"n = 353", 353, false},
		{"n = 354", 354, false},
		{"n = 355", 355, false},
		{"n = 356", 356, false},
		{"n = 357", 357, false},
		{"n = 358", 358, false},
		{"n = 359", 359, false},
		{"n = 360", 360, false},
		{"n = 361", 361, false},
		{"n = 362", 362, false},
		{"n = 363", 363, false},
		{"n = 364", 364, false},
		{"n = 365", 365, false},
		{"n = 366", 366, false},
		{"n = 367", 367, false},
		{"n = 368", 368, false},
		{"n = 369", 369, false},
		{"n = 370", 370, false},
		{"n = 371", 371, false},
		{"n = 372", 372, false},
		{"n = 373", 373, false},
		{"n = 374", 374, false},
		{"n = 375", 375, false},
		{"n = 376", 376, false},
		{"n = 377", 377, false},
		{"n = 378", 378, false},
		{"n = 379", 379, false},
		{"n = 380", 380, true},
		{"n = 381", 381, false},
		{"n = 382", 382, false},
		{"n = 383", 383, false},
		{"n = 384", 384, false},
		{"n = 385", 385, false},
		{"n = 386", 386, false},
		{"n = 387", 387, false},
		{"n = 388", 388, false},
		{"n = 389", 389, false},
		{"n = 390", 390, false},
		{"n = 391", 391, false},
		{"n = 392", 392, false},
		{"n = 393", 393, false},
		{"n = 394", 394, false},
		{"n = 395", 395, false},
		{"n = 396", 396, false},
		{"n = 397", 397, false},
		{"n = 398", 398, false},
		{"n = 399", 399, false},
		{"n = 400", 400, false},
		{"n = 2147441940", 2147441940, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := math.PronicNumber(test.n)
			if result != test.expectedValue {
				t.Errorf("expected value: %v, got: %v", test.expectedValue, result)
			}
		})
	}
}
func BenchmarkPronicNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.PronicNumber(65536)
	}
}
