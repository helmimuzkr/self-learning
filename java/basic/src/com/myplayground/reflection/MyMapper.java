package com.myplayground.reflection;

import java.lang.reflect.Field;

public class MyMapper {

    // Generic static Method
    public static <S, D> D objectMapper(S source, Class<D> destination) throws Exception{
        System.out.println("Mapper");

        // Get the field of Source class
        Field[] sourceField = source.getClass().getDeclaredFields();

        // Get the field of Destination class
        D newDestination = destination.newInstance(); // will throw exceptions
        Field[] newDestinationField = newDestination.getClass().getDeclaredFields();

        // Mapper operation
        for (int i = 0; i < sourceField.length; i++) {
            // Check if Source field has the same name and type as Destination field
            if (sourceField[i].getName().equals(newDestinationField[i].getName())
                    && sourceField[i].getType().equals(newDestinationField[i].getType())) {
                try {
                    // Set the modifier access Source field to public
                    // for get Source field and set the Destination field value
                    sourceField[i].setAccessible(true);
                    newDestinationField[i].setAccessible(true);
                    newDestinationField[i].set(newDestination, sourceField[i].get(source));
                } catch (Exception e) {
                    String fieldName = sourceField[i].getName();
                    String className = sourceField.getClass().getName();
                    throw new RuntimeException(String.format("failed to convert field %s from class %s", fieldName, className), e);
                }
            }
        }

        return newDestination;
    }
}
