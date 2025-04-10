package medium

import "math"

type Point struct {
	X, Y float64
}

// Big O: O(n)
// Não funciona, pois eu segui a seguinte linha de raciocínio:
// Estava separando o polígono em uma sequência triângulos e verificando se o ponto estava dentro de algum deles.
// Porém, o ponto pode estar dentro de algum dos triângulos, mas fora do polígono.
// Por exemplo, imagine um polígono com 5 pontos, formando 2 orelhas de gato, e imagine que o ponto a ser verificado está entre as duas orelhas.
//
//	   /\  .  /\
//	  /  \   /  \
//	 /    \ /    \
//	/_____________\
//
// Da maneira que eu fiz, o ponto estaria fora do polígono, pois a função iria chegar num triângulo
// que está aberto (para o lado de fora do polígono), e a função iria retornar que o ponto está dentro do polígono,
// quando na verdade ele não está, ele apenas está dentro de um triângulo que é formado por 3 pontos seguidos do polígono.
func isValidPointInPolygon(polygon []Point, point Point) bool {
	if len(polygon) < 3 {
		return false
	}

	for i := 0; i < len(polygon)-2; i++ {
		p1 := polygon[i]
		p2 := polygon[i+1]
		p3 := polygon[i+2]

		triangleArea := calculateTriangleArea(p1, p2, p3)

		areaWithPoint1 := calculateTriangleArea(point, p2, p3)
		areaWithPoint2 := calculateTriangleArea(point, p1, p2)
		areaWithPoint3 := calculateTriangleArea(point, p1, p3)

		if triangleArea == (areaWithPoint1 + areaWithPoint2 + areaWithPoint3) {
			return true
		}
	}

	return false
}

func calculateTriangleArea(p1, p2, p3 Point) float64 {
	return math.Abs(((p1.X * (p2.Y - p3.Y)) + (p2.X * (p3.Y - p1.Y)) + (p3.X * (p1.Y - p2.Y))) / 2)
}
