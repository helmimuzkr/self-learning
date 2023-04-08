package com.learnjava.simplerestapi.tck.internal;

import com.learnjava.simplerestapi.tck.domain.Title;
import org.springframework.data.jpa.repository.JpaRepository;

public interface ITitleRepository extends JpaRepository<Title, Long> {
}
