import numpy as np 
from sympy import *

def puntofijo():

    eqn_ = input("Introduce la ecuación f(x): ")
    eqn2_ = input("Introduce la función g(x): ")
    valorA_ = float(input("Introduce el valor inicial: "))
    tol_ = float(input("Introduce la tolerancia: "))

    eqn = sympify(eqn_)
    fx = lambda x:eqn.subs({'x':x})
    
    eqn2 = sympify(eqn2_)
    gx = lambda x:eqn2.subs({'x':x})
    
    a = valorA_
    tolera = tol_
    iteramax = 100000
    i = 0 
    b = gx(a)
    tramo = abs(b-a)

    while(tramo>=tolera and i<=iteramax):
        a = b
        b = gx(a)
        tramo = abs(b-a)
        i = i + 1
    respuesta = b
    
    if (i>=iteramax ):
        respuesta = np.nan
        
    return("la raiz es: " + str(respuesta), "Error: ", tramo)

resultado = puntofijo()
print(resultado)