-- CreateTable
CREATE TABLE "Screen" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "position" INTEGER NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "Screen_position_key" ON "Screen"("position");
