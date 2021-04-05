package main

import "fmt"

func main() {
	// 클로저 (Closure)
	// 익명함수 사용할 경우 함수를 변수에 할당해서 사용가능
	// 함수 안에서 함수를 선언 및 정의 가능 -> 이 때 외부 함수에 선언된 변수에 접근 가능한 함수
	// 함수가 선언되는 순간에 함수가 실행 될 때 실체의 외부 변수에 접근하기 위한 스냅샷(객체)
	// 함수를 호출 할 때 이전에 존재했던 값을 유지하기 위해서 -> 비동기, 누적카운트, 무분별한 전역변수 남발 ...
	// 남발하면 객체들이 메모리에 존재하므로, -> 메모리 부족, 오버플로 현상, 자원을 무분별하게 사용 가능성
	// 클로저를 정확하게 이해하고 사용

	// Example 1
	multiply := func(x, y int) int { // 익명함
		return x * y
	}

	r1 := multiply(5, 10)
	fmt.Println("Example 1 : ", r1)

	// Example 2
	m, n := 5, 10 // 지역변수 - 변수가 캡쳐
	sum := func(c int) int { // 익명함수 변수 할당
		return m + n + c // 지역 변수 소멸되지 않는다. (함수 호출 시 마다 사용 가능)
	}

	r2 := sum(10)
	fmt.Println("Example 2 :", r2)

}
