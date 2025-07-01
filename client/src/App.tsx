import {  Stack } from "@chakra-ui/react";
import { Container } from "@chakra-ui/react";
import  Navbar  from "./components/Navbar";
import  TodoForm  from "./components/TodoForm";
import  TodoList  from "./components/TodoList";
export default function App() {
  return (
    <Stack>
      <Navbar />
      <Container>
        <TodoForm/>
        <TodoList/>
      </Container>
    </Stack>
  );
}
