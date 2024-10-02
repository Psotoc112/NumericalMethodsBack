from sympy import sympify

def Newton():

    eqn_ = input("Introduce la ecuación f(x): ")
    eqn1_ = input("Introduce la derivada de f(x): ")
    xo = float(input("Introduce el valor inicial (xo): "))
    tol = float(input("Introduce la tolerancia (tol): "))
    

    eqn = sympify(eqn_)
    f = lambda x: eqn.subs({'x': x})
    
    eqn1 = sympify(eqn1_)
    g = lambda x: eqn1.subs({'x': x})
    

    niter = 100000
    cont = 0
    err = tol + 1
    

    while (err > tol) and (niter > cont):
        fxo = f(xo)
        gxo = g(xo)
        
        if gxo == 0:
            return "Derivada igual a cero en: " + str(cont) + " iteraciones"
        
        xn = xo - fxo / gxo
        err = abs(xn - xo)
        xo = xn
        cont += 1
    

    if err <= tol:
        return "Se encontró una raíz en: " + str(xo) + "\nCon: " + str(cont) + " iteraciones"
    else:
        return "El método no logró converger"

# Ejecutar la función
resultado = Newton()
print(resultado)
