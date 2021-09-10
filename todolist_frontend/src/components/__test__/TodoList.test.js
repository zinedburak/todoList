import React, { useState } from "react"
import {render,fireEvent, queryByPlaceholderText} from '@testing-library/react'
import "@testing-library/jest-dom/extend-expect"
import TodoList from "../TodoList"



it("renders  with todos", () => {
    const todos = [{text: "Dummy Test",id: 0},{text: "Dummy Test2", id: 1}]
    const component = render(<TodoList todos={todos}/>)


    const genral_form = component.getByTestId("todoList");
    const todos_ex = component.getAllByTestId("todo");
    const todos_count = component.getAllByTestId("todo").length;
    expect(genral_form).toBeTruthy;
    expect(todos_ex).toBeTruthy;
    expect(todos_count === 2).toBeTruthy;
    

});

it("renders  without todos", () => {
    const todos = []
    const component = render(<TodoList todos={todos}/>)


    const genral_form = component.getByTestId("todoList");
    const single_todo = component.queryByTestId("todo");
    expect(genral_form).toBeTruthy;
    expect(single_todo).toBeNull;
    

});