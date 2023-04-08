package com.learnjava.simplerestapi.tck.internal;

import com.learnjava.simplerestapi.tck.domain.Title;

import java.util.List;

public interface ITitleService {
    List<Title> getTitle();
    Title getTitleById(int id);
    void addTitle(Title title);
    void updateTitle(int id, Title title);
    void deleteTitle(int id);
}
