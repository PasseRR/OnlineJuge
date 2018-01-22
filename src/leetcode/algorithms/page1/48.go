package page1

func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length-1; i++ {
		//i只能取到 length - 2, 因为length - 1是对称轴
		for j := 0; j < length-1-i; j++ {
			//j只能取到length - 1 - i, 在对称轴的左边
			matrix[i][j], matrix[length-1-j][length-1-i] = matrix[length-1-j][length-1-i], matrix[i][j]
		}
	}
	for i := 0; i < length/2; i++ {
		//i只能取到横向中间轴的上面
		for j := 0; j < length; j++ {
			//j可以取到所有值
			//按横向轴翻转, j不变；i变为length-1-i
			matrix[i][j], matrix[length-1-i][j] = matrix[length-1-i][j], matrix[i][j]
		}
	}
}
