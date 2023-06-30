package com.myplayground.reflection;

import com.myplayground.reflection.data.Book;
import com.myplayground.reflection.data.DtoBook;
import com.myplayground.reflection.data.DtoUser;
import com.myplayground.reflection.data.User;

public class Main {

    public static void main(String[] args) {

        DtoUser dtoUser = new DtoUser("helmi", "123");

        // Simple example
        ClassReflection.classReflectionExample(dtoUser);
        FieldReflection.fieldReflectionExample(dtoUser);
        FieldReflection.changeAccessModifierExample(dtoUser);
        MethodReflection.methodReflectionExample(dtoUser);

        // Advance Example : Object Mapper
        try {
            User domainUser = MyMapper.objectMapper(dtoUser, User.class);
            System.out.println(domainUser);

            // Field akan tidak sesuai dengan ekspetasi.
            // Karena value akan di mapping sesuai urutan Field nya juga.
            // Solusi: harus menggunakan Alogirtma Selection. (PR)
            DtoBook dtoBook = new DtoBook("tokek");
            Book domainBook = MyMapper.objectMapper(dtoBook, Book.class);
            System.out.println(domainBook);
        }catch (Exception e) {
            System.out.println(e);
        }
    }
}
