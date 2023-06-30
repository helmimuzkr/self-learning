package com.myplayground.reflection;

import com.myplayground.reflection.data.DtoUser;

import java.lang.reflect.Method;
import java.util.Arrays;
import java.util.Optional;

public class MethodReflection {

    static void methodReflectionExample(DtoUser user) {
        System.out.println("Get Declared Methods");

        Method[] declaredMethods = user.getClass().getDeclaredMethods();
        Arrays.stream(declaredMethods).forEach(System.out::println);

        Optional<Method> firstMethod = Arrays.stream(declaredMethods).findFirst();
        System.out.println("return type in first method -> " + firstMethod.get().getReturnType());
        System.out.println("parameter type in first method -> " + Arrays.toString(firstMethod.get().getParameterTypes()));
        System.out.println("exception in first method -> " + Arrays.toString(firstMethod.get().getExceptionTypes()));

        System.out.println();
    }
}
