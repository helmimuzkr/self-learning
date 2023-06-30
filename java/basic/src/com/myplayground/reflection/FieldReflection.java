package com.myplayground.reflection;

import com.myplayground.reflection.data.DtoUser;

import java.lang.reflect.Field;
import java.util.Arrays;

public class FieldReflection {

    static void fieldReflectionExample(DtoUser user) {
        System.out.println("Get Fields");
        // Get fields -> method is used to obtain the public fields of a class and its superclasses.
        Field[] fields = user.getClass().getFields();
        Arrays.stream(fields).forEach(System.out::println);
        System.out.println();

        System.out.println("Get Declared Fields");
        // Get declared fields ->  method is used to obtain a specific field of a class, regardless of its accessibility.
        // It considers only the fields that are directly declared in the class, not the fields inherited from its superclasses.
        Field[] declaredFields = user.getClass().getDeclaredFields();
        Arrays.stream(declaredFields).forEach(System.out::println);
        System.out.println();
    }

    static void changeAccessModifierExample(DtoUser user) {
        String oldUsername = user.getUsername();

        System.out.println("Access Private Field");
        try {
            Field field = user.getClass().getDeclaredField("username");
            // change value of private field
            field.setAccessible(true);
            field.set(user, "Muzakir"); // from helmi to Muzakir
            String username = field.get(user).toString();
            System.out.println("old username: " + oldUsername);
            System.out.println("new username: " + username);
        }catch (Exception e) {System.out.println(e);}

        System.out.println();
    }
}
