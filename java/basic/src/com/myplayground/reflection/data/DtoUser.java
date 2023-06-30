package com.myplayground.reflection.data;

public class DtoUser {
    private String username;
    private String password;
    public boolean isActive;

    public DtoUser() {}

    public DtoUser(String username, String password) {
        this.username = username;
        this.password = password;
    }

    public String getUsername() {
        return username;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}