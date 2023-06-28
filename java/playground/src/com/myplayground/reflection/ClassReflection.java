package com.myplayground.reflection;

import com.myplayground.reflection.data.DtoUser;

public class ClassReflection {

    static void classReflectionExample(DtoUser user) {
        System.out.println("Get Class Name");

        Class<?> userClass = user.getClass(); // orr DtoUser.class
        String className = userClass.getName();
        System.out.println(className);

        System.out.println();
    }
}
