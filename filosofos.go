package main
import "fmt"
var filosofosSaciados int = 1
type Tenedor struct{
	tomado int
}
var Tenedores[5] Tenedor

type filosofo struct{
	Izquierda int
	Derecha int
}
var todosfilosofos[5] filosofo

func cara(){
	var cantidad int
	fmt.Println("ingrese cantidad de filosofos : ")
	fmt.Scan(&cantidad)
}

func Cenar(recorreF int){
	var n int = 5
    if todosfilosofos[recorreF].Izquierda==3 && todosfilosofos[recorreF].Derecha==3{
		fmt.Println("\tFilosofo ", (recorreF+1)," Termino su cena\n")
	}
    if todosfilosofos[recorreF].Izquierda==1 && todosfilosofos[recorreF].Derecha==1{

		fmt.Println("\tFilosofo ",(recorreF+1)," Termino su cena\n")

		todosfilosofos[recorreF].Izquierda = 3
		todosfilosofos[recorreF].Derecha = 3
		var otrotenedor int = recorreF-1
		if otrotenedor == (-1) {
			otrotenedor=(n-1);
		}
		Tenedores[recorreF].tomado = 0
		Tenedores[otrotenedor].tomado = 0
		fmt.Println("\tFilosofo ",(recorreF+1)," dejo tenedor ",(recorreF+1)," y tenedor ",(otrotenedor+1),"\n")
		filosofosSaciados++;
    }
	if todosfilosofos[recorreF].Izquierda==1 && todosfilosofos[recorreF].Derecha==0 {
		if recorreF==(n-1) {
			if Tenedores[recorreF].tomado==0 {
				Tenedores[recorreF].tomado = 1
				todosfilosofos[recorreF].Derecha = 1
				fmt.Println("\ttenedor ",(recorreF+1)," tomado por Filosofo ",(recorreF+1),"\n")
			}else{
				fmt.Println("\tFilosofo ",(recorreF+1)," esta esperando para tenedor ",(recorreF+1),"\n")
			}
		}else{
			var IDdublicado int = recorreF ;
			recorreF = recorreF-1;

			if recorreF== -1{
				recorreF =(n-1);
			}
			if Tenedores[recorreF].tomado == 0{
				Tenedores[recorreF].tomado = 1
				todosfilosofos[IDdublicado].Derecha = 1
				fmt.Println("\ttenedor ",(recorreF+1)," tomado por Filosofo ",(IDdublicado+1),"\n")
			}else{
				fmt.Println("\tFilosofo ",(IDdublicado+1)," esta esperando por el tenedor ",(recorreF+1),"\n")
			}
		}
	}
	if todosfilosofos[recorreF].Izquierda==0 {
		if recorreF==(n-1) {
			if Tenedores[recorreF-1].tomado == 0 {
				Tenedores[recorreF-1].tomado = 1
				todosfilosofos[recorreF].Izquierda = 1
				fmt.Println("\ttenedor ",recorreF," tomado por Filosofo ",(recorreF+1),"\n")
			}else{
				fmt.Println("\tFilosofo ",(recorreF+1)," esta esperando por el tenedor ",recorreF,"\n")
			}
		}else{
			if Tenedores[recorreF].tomado == 0{
				Tenedores[recorreF].tomado = 1
				todosfilosofos[recorreF].Izquierda = 1
				fmt.Println("\ttenedor ",(recorreF+1)," tomado por Filosofo ",(recorreF+1),"\n")
			}else{
				fmt.Println("\tFilosofo ",(recorreF+1)," esta esperando por el tenedor ",(recorreF+1),"\n")
			}
		}
	}

}
func main(){
	var n int = 5
	for i:=0; i<n ;i++{
		Tenedores[i].tomado = 0
		todosfilosofos[i].Izquierda = 0
		todosfilosofos[i].Derecha = 0
	}
	for filosofosSaciados<n{
		for i:=0;i<n;i++ {
			Cenar(i);
		}
		fmt.Println("\t hasta ahora el numero del filosofo que Termino la cena es ",filosofosSaciados,"\n\n")
	}
}
