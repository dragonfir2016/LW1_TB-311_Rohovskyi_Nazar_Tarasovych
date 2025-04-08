package main

import (
	"fmt"
	"math"
)

func round(x float64) float64 {
	return math.Round(x*100) / 100
}

func task1() {
	var H_p, C_p, S_p, N_p, O_p, W_p, A_p float64

	fmt.Println("Введіть значення для завдання 1:")
	fmt.Print("H_p, C_p, S_p, N_p, O_p, W_p, A_p (через пробіл): ")
	_, err := fmt.Scanln(&H_p, &C_p, &S_p, &N_p, &O_p, &W_p, &A_p)
	if err != nil {
		fmt.Println("Помилка введення:", err)
		return
	}

	K_PC := 100 / (100 - W_p)
	K_PT := 100 / (100 - W_p - A_p)

	H_C := H_p * K_PC
	C_C := C_p * K_PC
	S_C := S_p * K_PC
	N_C := N_p * K_PC
	O_C := O_p * K_PC
	A_C := A_p * K_PC

	H_T := H_p * K_PT
	C_T := C_p * K_PT
	S_T := S_p * K_PT
	N_T := N_p * K_PT
	O_T := O_p * K_PT

	Q_P_H := 339*C_p + 1030*H_p - 108.8*(O_p-S_p) - 25*W_p
	Q_C_H := (Q_P_H + 0.025*W_p) * (100 / (100 - W_p))
	Q_T_H := (Q_P_H + 0.025*W_p - A_p) * (100 / (100 - W_p - A_p))

	fmt.Println("\nКоефіцієнти переходу:")
	fmt.Printf("K_PC: %.2f\n", round(K_PC))
	fmt.Printf("K_PT: %.2f\n", round(K_PT))
	fmt.Println("\nСклад сухої маси:")
	fmt.Printf("H_C: %.2f\n", round(H_C))
	fmt.Printf("C_C: %.2f\n", round(C_C))
	fmt.Printf("S_C: %.2f\n", round(S_C))
	fmt.Printf("N_C: %.2f\n", round(N_C))
	fmt.Printf("O_C: %.2f\n", round(O_C))
	fmt.Printf("A_C: %.2f\n", round(A_C))
	fmt.Println("\nСклад горючої маси:")
	fmt.Printf("H_T: %.2f\n", round(H_T))
	fmt.Printf("C_T: %.2f\n", round(C_T))
	fmt.Printf("S_T: %.2f\n", round(S_T))
	fmt.Printf("N_T: %.2f\n", round(N_T))
	fmt.Printf("O_T: %.2f\n", round(O_T))
	fmt.Println("\nНижча теплота згоряння:")
	fmt.Printf("Q_P_H: %.2f кДж/кг\n", round(Q_P_H))
	fmt.Printf("Q_C_H: %.2f кДж/кг\n", round(Q_C_H))
	fmt.Printf("Q_T_H: %.2f кДж/кг\n", round(Q_T_H))
}

func task2() {
	var H_g, C_g, S_g, O_g, A_g, W_p_mazut, Q_g float64

	fmt.Println("Введіть значення для завдання 2:")
	fmt.Print("H_g, C_g, S_g, O_g, A_g, W_p_mazut, Q_g (через пробіл): ")
	_, err := fmt.Scanln(&H_g, &C_g, &S_g, &O_g, &A_g, &W_p_mazut, &Q_g)
	if err != nil {
		fmt.Println("Помилка введення:", err)
		return
	}

	K_GC := 100 / (100 - W_p_mazut - A_g)

	H_p := H_g * K_GC
	C_p := C_g * K_GC
	S_p := S_g * K_GC
	O_p := O_g * K_GC
	A_p := A_g * (100 / (100 - W_p_mazut))

	Q_p := Q_g * (100 - W_p_mazut - A_g) / 100

	fmt.Println("\nКоефіцієнт переходу до робочої маси:")
	fmt.Printf("K_GC: %.2f\n", round(K_GC))
	fmt.Println("\nСклад робочої маси:")
	fmt.Printf("H_p: %.2f\n", round(H_p))
	fmt.Printf("C_p: %.2f\n", round(C_p))
	fmt.Printf("S_p: %.2f\n", round(S_p))
	fmt.Printf("O_p: %.2f\n", round(O_p))
	fmt.Printf("A_p: %.2f\n", round(A_p))
	fmt.Println("\nНижча теплота згоряння:")
	fmt.Printf("Q_p: %.2f МДж/кг\n", round(Q_p))
}

func main() {
	var choice int
	fmt.Println("Виберіть завдання:")
	fmt.Println("1 - Завдання 1")
	fmt.Println("2 - Завдання 2")
	fmt.Print("Ваш вибір: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Помилка введення:", err)
		return
	}

	switch choice {
	case 1:
		task1()
	case 2:
		task2()
	default:
		fmt.Println("Невірний вибір. Спробуйте ще раз.")
	}
}
