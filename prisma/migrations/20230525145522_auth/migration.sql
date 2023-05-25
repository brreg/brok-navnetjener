-- CreateTable
CREATE TABLE "Shareholder" (
    "id" SERIAL NOT NULL,
    "name" TEXT NOT NULL,
    "sosialSecurityNumber" BIGINT NOT NULL,
    "ethereumWallet" TEXT NOT NULL,

    CONSTRAINT "Shareholder_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "Shareholder_ethereumWallet_key" ON "Shareholder"("ethereumWallet");
