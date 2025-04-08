package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

func round(x float64) float64 {
	return math.Round(x*100) / 100
}

var indexTmpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <title>Лабораторна робота №2</title>
</head>
<body>
    <h1>Лабораторна робота №2</h1>
    <ul>
        <li><a href="/task1">Завдання 1</a></li>
        <li><a href="/task2">Завдання 2</a></li>
    </ul>
</body>
</html>
`))

var task1Tmpl = template.Must(template.New("task1").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <title>Завдання 1</title>
</head>
<body>
    <h1>Завдання 1</h1>
    <form method="post" action="/task1">
        <p>Введіть значення:</p>
        <p>H_p: <input type="text" name="H_p"></p>
        <p>C_p: <input type="text" name="C_p"></p>
        <p>S_p: <input type="text" name="S_p"></p>
        <p>N_p: <input type="text" name="N_p"></p>
        <p>O_p: <input type="text" name="O_p"></p>
        <p>W_p: <input type="text" name="W_p"></p>
        <p>A_p: <input type="text" name="A_p"></p>
        <input type="submit" value="Розрахувати">
    </form>
    {{if .}}
    <h2>Результати розрахунків</h2>
    <h3>Коефіцієнти переходу:</h3>
    <p>K_PC: {{printf "%.2f" .K_PC}}</p>
    <p>K_PT: {{printf "%.2f" .K_PT}}</p>
    <h3>Склад сухої маси:</h3>
    <p>H_C: {{printf "%.2f" .H_C}}</p>
    <p>C_C: {{printf "%.2f" .C_C}}</p>
    <p>S_C: {{printf "%.2f" .S_C}}</p>
    <p>N_C: {{printf "%.2f" .N_C}}</p>
    <p>O_C: {{printf "%.2f" .O_C}}</p>
    <p>A_C: {{printf "%.2f" .A_C}}</p>
    <h3>Склад горючої маси:</h3>
    <p>H_T: {{printf "%.2f" .H_T}}</p>
    <p>C_T: {{printf "%.2f" .C_T}}</p>
    <p>S_T: {{printf "%.2f" .S_T}}</p>
    <p>N_T: {{printf "%.2f" .N_T}}</p>
    <p>O_T: {{printf "%.2f" .O_T}}</p>
    <h3>Нижча теплота згоряння:</h3>
    <p>Q_P_H: {{printf "%.2f" .Q_P_H}} кДж/кг</p>
    <p>Q_C_H: {{printf "%.2f" .Q_C_H}} кДж/кг</p>
    <p>Q_T_H: {{printf "%.2f" .Q_T_H}} кДж/кг</p>
    {{end}}
    <p><a href="/">На головну</a></p>
</body>
</html>
`))

var task2Tmpl = template.Must(template.New("task2").Parse(`
<!DOCTYPE html>
<html lang="uk">
<head>
    <meta charset="UTF-8">
    <title>Завдання 2</title>
</head>
<body>
    <h1>Завдання 2</h1>
    <form method="post" action="/task2">
        <p>Введіть значення:</p>
        <p>H_g: <input type="text" name="H_g"></p>
        <p>C_g: <input type="text" name="C_g"></p>
        <p>S_g: <input type="text" name="S_g"></p>
        <p>O_g: <input type="text" name="O_g"></p>
        <p>A_g: <input type="text" name="A_g"></p>
        <p>W_p_mazut: <input type="text" name="W_p_mazut"></p>
        <p>Q_g: <input type="text" name="Q_g"></p>
        <input type="submit" value="Розрахувати">
    </form>
    {{if .}}
    <h2>Результати розрахунків</h2>
    <h3>Коефіцієнт переходу до робочої маси:</h3>
    <p>K_GC: {{printf "%.2f" .K_GC}}</p>
    <h3>Склад робочої маси:</h3>
    <p>H_p: {{printf "%.2f" .H_p}}</p>
    <p>C_p: {{printf "%.2f" .C_p}}</p>
    <p>S_p: {{printf "%.2f" .S_p}}</p>
    <p>O_p: {{printf "%.2f" .O_p}}</p>
    <p>A_p: {{printf "%.2f" .A_p}}</p>
    <h3>Нижча теплота згоряння:</h3>
    <p>Q_p: {{printf "%.2f" .Q_p}} МДж/кг</p>
    {{end}}
    <p><a href="/">На головну</a></p>
</body>
</html>
`))

type Task1Data struct {
	K_PC  float64
	K_PT  float64
	H_C   float64
	C_C   float64
	S_C   float64
	N_C   float64
	O_C   float64
	A_C   float64
	H_T   float64
	C_T   float64
	S_T   float64
	N_T   float64
	O_T   float64
	Q_P_H float64
	Q_C_H float64
	Q_T_H float64
}

type Task2Data struct {
	K_GC float64
	H_p  float64
	C_p  float64
	S_p  float64
	O_p  float64
	A_p  float64
	Q_p  float64
}

func task1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		H_p, err1 := strconv.ParseFloat(r.FormValue("H_p"), 64)
		C_p, err2 := strconv.ParseFloat(r.FormValue("C_p"), 64)
		S_p, err3 := strconv.ParseFloat(r.FormValue("S_p"), 64)
		N_p, err4 := strconv.ParseFloat(r.FormValue("N_p"), 64)
		O_p, err5 := strconv.ParseFloat(r.FormValue("O_p"), 64)
		W_p, err6 := strconv.ParseFloat(r.FormValue("W_p"), 64)
		A_p, err7 := strconv.ParseFloat(r.FormValue("A_p"), 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil {
			http.Error(w, "Невірні вхідні дані. Перевірте введені значення.", http.StatusBadRequest)
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

		data := Task1Data{
			K_PC:  round(K_PC),
			K_PT:  round(K_PT),
			H_C:   round(H_C),
			C_C:   round(C_C),
			S_C:   round(S_C),
			N_C:   round(N_C),
			O_C:   round(O_C),
			A_C:   round(A_C),
			H_T:   round(H_T),
			C_T:   round(C_T),
			S_T:   round(S_T),
			N_T:   round(N_T),
			O_T:   round(O_T),
			Q_P_H: round(Q_P_H),
			Q_C_H: round(Q_C_H),
			Q_T_H: round(Q_T_H),
		}
		task1Tmpl.Execute(w, data)
		return
	}
	task1Tmpl.Execute(w, nil)
}

func task2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		H_g, err1 := strconv.ParseFloat(r.FormValue("H_g"), 64)
		C_g, err2 := strconv.ParseFloat(r.FormValue("C_g"), 64)
		S_g, err3 := strconv.ParseFloat(r.FormValue("S_g"), 64)
		O_g, err4 := strconv.ParseFloat(r.FormValue("O_g"), 64)
		A_g, err5 := strconv.ParseFloat(r.FormValue("A_g"), 64)
		W_p_mazut, err6 := strconv.ParseFloat(r.FormValue("W_p_mazut"), 64)
		Q_g, err7 := strconv.ParseFloat(r.FormValue("Q_g"), 64)
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil {
			http.Error(w, "Невірні вхідні дані. Перевірте введені значення.", http.StatusBadRequest)
			return
		}

		K_GC := 100 / (100 - W_p_mazut - A_g)
		H_p := H_g * K_GC
		C_p := C_g * K_GC
		S_p := S_g * K_GC
		O_p := O_g * K_GC
		A_p := A_g * (100 / (100 - W_p_mazut))
		Q_p := Q_g * (100 - W_p_mazut - A_g) / 100

		data := Task2Data{
			K_GC: round(K_GC),
			H_p:  round(H_p),
			C_p:  round(C_p),
			S_p:  round(S_p),
			O_p:  round(O_p),
			A_p:  round(A_p),
			Q_p:  round(Q_p),
		}
		task2Tmpl.Execute(w, data)
		return
	}
	task2Tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/task1", task1Handler)
	http.HandleFunc("/task2", task2Handler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTmpl.Execute(w, nil)
	})
	fmt.Println("Сервер запущено на порті 8080...")
	http.ListenAndServe(":8080", nil)
}
