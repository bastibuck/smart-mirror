import { TodoistApi } from "@doist/todoist-api-typescript";

import { env } from "~/env";

import { createTRPCRouter, publicProcedure } from "~/server/api/trpc";

const todoist = new TodoistApi(env.TODOIST_API_TOKEN);

export const todoRouter = createTRPCRouter({
  due: publicProcedure.query(async () => {
    const dueToDos = (await todoist.getTasks({ filter: "today | overdue" }))
      .sort((a, b) => {
        if (!a.due?.date || !b.due?.date) {
          return 0;
        }

        return a.due.date.localeCompare(b.due.date);
      })
      .map((todo) => {
        const dueDate = new Date(todo.due?.date ?? new Date());
        dueDate.setHours(23, 59, 59, 999);

        return {
          ...todo,
          isOverdue: dueDate.getTime() < new Date().getTime(),
        };
      });

    return dueToDos;
  }),
});
