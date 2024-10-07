import numpy as np
import math

from math import sin, cos, exp, log

def busquedas_incrementales(f, x0, intervalo, tol, max_iter):
    """
    Método de Búsquedas Incrementales 
    """
    # Inicializar variables
    x_i = x0
    iteraciones = 0
    tabla = []  # Almacenar las iteraciones en una tabla

    # Evaluar la función y capturar errores de dominio
    try:
        f_xi = f(x_i)
    except (ValueError, ZeroDivisionError) as e:
        print(f"Error de dominio al evaluar f({x_i}): {e}")
        return None

    # Recorrer con pasos de tamaño del intervalo mientras no exceda el número de iteraciones
    while iteraciones < max_iter:
        # Calcular el siguiente punto
        x_i_next = x_i + intervalo
        try:
            f_xi_next = f(x_i_next)
        except (ValueError, ZeroDivisionError) as e:
            print(f"Error de dominio al evaluar f({x_i_next}): {e}")
            return None

        # Calcular error absoluto 
        error_abs = abs(x_i_next - x_i)

        # Guardar valores en la tabla
        tabla.append([iteraciones + 1, x_i, f_xi, x_i_next, f_xi_next, error_abs])

        # Verificar si hay un cambio de signo 
        if f_xi * f_xi_next < 0:
            print(f"Intervalo con raíz aproximada: [{x_i:.4f}, {x_i_next:.4f}]\n")
            return x_i, x_i_next  

        # Verificar tolerancia
        if error_abs < tol:
            print(f"El método ha alcanzado la tolerancia especificada Tol = {tol} en la iteración {iteraciones + 1}.")
            return x_i, x_i_next

        # Avanzar al siguiente punto
        x_i = x_i_next
        f_xi = f_xi_next
        iteraciones += 1

    print(f"No se encontraron raíces en el intervalo dado después de {max_iter} iteraciones.")
    return None

# Función para procesar la funcion para las librerias
def procesar_funcion(funcion_str):

    # Reemplazos necesarios para asegurar compatibilidad
    funcion_str = funcion_str.replace("ln(", "log(")  # Cambiar ln a log para logaritmo natural
    funcion_str = funcion_str.replace("^", "**")       # Cambiar ^ a ** para potencias
    return funcion_str

# Función para ingresar los parámetros
def ingresar_datos():

    # Solicitar la función y realizar las conversiones necesarias
    funcion_str = input("Ingrese la función f(x) (por ejemplo, ln(sin(x)**2 + 1) - 1/2): ")
    funcion_str = procesar_funcion(funcion_str)
    
    # Convertir la función en un objeto evaluable
    f = lambda x: eval(funcion_str, {"x": x, "sin": sin, "cos": cos, "exp": exp, "log": log, "math": math})
    
    # Solicitar el valor inicial
    x0 = float(input("Ingrese el valor inicial x0: "))

    # Solicitar el tamaño del intervalo
    intervalo = float(input("Ingrese el tamaño del intervalo: "))

    # Solicitar la tolerancia
    tol = float(input("Ingrese la tolerancia Tol (por ejemplo, 1e-7): "))

    # Solicitar el número máximo de iteraciones
    max_iter = int(input("Ingrese el número máximo de iteraciones N: "))

    return f, x0, intervalo, tol, max_iter

# Ingreso de los parámetros y la función
funcion, x0, intervalo, tol, max_iter = ingresar_datos()

# Ejecutar el método 
busquedas_incrementales(funcion, x0, intervalo, tol, max_iter)
