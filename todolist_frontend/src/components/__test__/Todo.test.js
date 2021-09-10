import React from "react"
import { render } from '@testing-library/react'
import Todo from "../Todo"
import "@testing-library/jest-dom/extend-expect"


it("renders correctly", () => {
    const todo = {text: "Dummy Test",id: 0}
    const component = render(<Todo todo={todo} />)

    const todo_item = component.getByTestId("todo")
    const todo_text = component.getByText(/Dummy Test/)
    expect(todo_item).toBeTruthy;
    expect(todo_text).toBeTruthy;
    

});