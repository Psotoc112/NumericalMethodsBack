import numpy as np

# Implementación del método de pivoteo parcial
def pivoteo_parcial(A, b):
    # Unión de A y b para formar la matriz aumentada
    matriz_aumentada = np.hstack((A, b.reshape(-1, 1))).astype(float)

    print("Matriz Aumentada:\n", matriz_aumentada, "\n")

    # Aplicación del método de pivoteo parcial
    n = len(A)
    for i in range(n):
        # Seleccionar el valor absoluto más grande en la columna i, desde la fila i hacia abajo
        max_index = np.argmax(abs(matriz_aumentada[i:, i])) + i

        # Intercambiar la fila actual con la fila que tiene el valor absoluto mayor
        if max_index != i:
            matriz_aumentada[[i, max_index]] = matriz_aumentada[[max_index, i]]
            print(f"Intercambio de fila {i + 1} con fila {max_index + 1} con pivoteo parcial.\n")

        # Transformar en cero las entradas de la columna i en las filas debajo del pivote
        for j in range(i + 1, n):
            factor = matriz_aumentada[j, i] / matriz_aumentada[i, i]
            matriz_aumentada[j, i:] -= factor * matriz_aumentada[i, i:]

        # Mostrar la matriz intermedia después de cada paso de eliminación
        print(f"Matriz intermedia después de la eliminación en la columna {i + 1}:\n", matriz_aumentada, "\n")

    print("Matriz Triangular Superior:\n", matriz_aumentada, "\n")

    # Sustitución regresiva para encontrar las soluciones
    x = np.zeros(n)
    for i in range(n - 1, -1, -1):
        x[i] = (matriz_aumentada[i, -1] - np.dot(matriz_aumentada[i, i + 1:n], x[i + 1:n])) / matriz_aumentada[i, i]

    print("Soluciones finales del sistema: ")
    for i in range(n):
        print(f"x{i + 1} = {x[i]:.4f}")
    return x

# Función para ingresar la matriz 
def ingresar_matriz_y_vector():
    """
    Solicitar al usuario los valores de la matriz A y del vector b
    """
    n = int(input("Ingrese el número de filas y columnas (matriz cuadrada): "))

    # Ingreso de la matriz A
    A = np.zeros((n, n))
    for i in range(n):
        while True:
            fila = input(f"Ingrese los elementos de la fila {i + 1} separados por espacio: ").split()
            if len(fila) == n:  # Verificar si la cantidad de elementos es correcta
                A[i] = [float(x) for x in fila]
                break
            else:
                print(f"Error: Debe ingresar exactamente {n} elementos en la fila {i + 1}. Inténtelo nuevamente.")

    # Ingreso del vector b
    print(f"\nIngrese el vector de términos independientes (b) separados por espacio: ")
    while True:
        elementos_b = input().split()
        if len(elementos_b) == n:  # Verificar si la cantidad de elementos es correcta
            b = np.array([float(x) for x in elementos_b])
            break
        else:
            print(f"Error: Debe ingresar exactamente {n} elementos en el vector b. Inténtelo nuevamente.")

    return A, b

# Ingreso de la matriz y vector
A_matriz, b_vector = ingresar_matriz_y_vector()

# Ejecutar el método
resultado = pivoteo_parcial(A_matriz, b_vector)
