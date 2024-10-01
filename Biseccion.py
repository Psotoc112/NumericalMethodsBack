from sympy import *
def biseccion():

    eqn = input("Introduce la ecuación f(x): ")
    xi = float(input("Introduce el valor inicial: "))
    xf = float(input("Introduce el valor final: "))
    tol = float(input("Introduce la tolerancia (tol): "))


    eqn = sympify(eqn)
    f = lambda x:eqn.subs({'x':x})
    niter = 100000
    xm = 0
    if(f(xi)*f(xf) == 0):
        if(f(xi) == 0):
            return("Xi es raíz")
        else:
            return("Xf es raíz")
    elif(f(xi)*f(xf) > 0):
        return("No se puede asegurar una raíz")
    else:
        xm = (xf+xi)/2
        cont = 0
        err = abs(xi - xf)
        while(err > tol and niter > cont and f(xm) != 0):
            if(f(xi)*f(xm) < 0):
                xf = xm
            else:
                xi = xm
            xm = (xf+xi)/2
            err = abs(xm - xi)
            cont= cont + 1
        if(f(xm) == 0):
           return("xm es raíz: " + str(xm) +"\n Iteraciones realizadas: " + str(cont))
        elif(err < tol):
            return("xm es raíz con tolerancia: " + str(tol) + " con xm: " + str(xm) +"\nIteraciones realizadas: " + str(cont))
        else:
            return("No se obtuvo solución")

resultado = biseccion()
print(resultado)