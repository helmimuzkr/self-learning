package com.playground.springcore.randomthings;

public class ClassUtil {

    public static <T> String getClassName(T instance) {
        return "Class Name -> " + instance.getClass().getName();
    }

    public static <T> T getInstance(Class<T> inputClass) throws IllegalAccessException, InstantiationException {
        return inputClass.newInstance();
    }
}
