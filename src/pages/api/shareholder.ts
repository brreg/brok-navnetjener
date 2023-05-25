// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next';
import debug from 'debug';
import { ApiError } from 'next/dist/server/api-utils';
import { PrismaClient } from '@prisma/client';
import { ethers } from 'ethers';

const log = debug('navnetjener:api:shareholder');

const prisma = new PrismaClient();

export default async function handler(req: NextApiRequest, res: NextApiResponse) {
  log(`HTTP ${req.method} ${req.url}`);
  try {
    switch (req.method) {
      case 'GET': {
        const { ethereumWallet } = parseQuery(req.query);
        log(`query: ethereumWallet ${ethereumWallet}`)
        const shareholder = await prisma.shareholder.findUnique({
          where: {
            ethereumWallet,
          },
          select: {
            name: true,
            ethereumWallet: true,
          },
        });
        log(`HTTP Response 200, shareholder ${JSON.stringify(shareholder)}`)
        return res.status(200).json({shareholder: shareholder});
      }

      case 'POST': {
        const { ethereumWallet, sosialSecurityNumber, name } = parseBody(req.body)
        log(`body: ethereumWallet ${ethereumWallet}, sosialSecurityNumber ${sosialSecurityNumber}, name ${name}`)
        const shareholder = await prisma.shareholder.findUnique({
          where: {
            ethereumWallet,
          },
          select: {
            name: true,
            ethereumWallet: true,
          },
        });
        if (shareholder) {
          log(`HTTP Response 200, shareholder already exist ${JSON.stringify(shareholder)}`)
          return res.status(200).json({shareholder: shareholder});
        }

        const newShareholder = await prisma.shareholder.create({
          data: {
            name,
            sosialSecurityNumber,
            ethereumWallet
          },
          select: {
            name: true,
            ethereumWallet: true,
          }
        })
        log(`HTTP Response 200, shareholder ${JSON.stringify(newShareholder)}`)
        return res.status(200).json({shareholder: newShareholder});
      }
      default: {
        log(`HTTP Response 405, Method ${req.method} Not Allowed`)
        res.setHeader('Allow', ['GET', 'POST']);
        res.status(405).end(`Method ${req.method} Not Allowed`);
      }
    }
  } catch (error) {
    if (error instanceof ApiError) {
      log(`HTTP Response ${error.statusCode}, ${error.message} ${error}`)
      return res.status(error.statusCode).json({ 
        status: error.statusCode,
        message: error.message,
      })
    }
    log(`HTTP Response 500, error ${error}`)
    return res.status(500).send(error)
  }
}

function parseBody(body: any) {
  if (!('ethereumWallet' in body)) {
    log("test")
    throw new ApiError(400, 'No ethereumWallet provided in body');
  }
  if (!('sosialSecurityNumber' in body)) {
    throw new ApiError(400, 'No sosialSecurityNumber provided in body');
  }
  if (!('name' in body)) {
    throw new ApiError(400, 'No name provided in body');
  }
  if (body.sosialSecurityNumber.length !== 11) {
		throw new ApiError(400, `sosialSecurityNumber ${body.sosialSecurityNumber} must be eleven digits`);
	}

  const sosialSecurityNumber : number = parseInt(body.sosialSecurityNumber)
  const name : string = body.name.toString()

  // ---
  const ethereumWallet = body.ethereumWallet.toString()

  // let ethereumWallet: string | undefined = undefined;
  // try {
  //   ethereumWallet = ethers.getAddress(body.ethereumWallet);
  // } catch (error) {
  //   throw new ApiError(400, `${body.ethereumWallet} is not an valid ethereum ethereumWallet`);
  // }
  // if (!ethereumWallet) {
  //   throw new ApiError(400, `Unknown error while parsing ${body.ethereumWallet} as an ethereum ethereumWallet`);
  // }

  // return { ethereumWallet, sosialSecurityNumber, name };
  
  return { ethereumWallet, sosialSecurityNumber, name };

}

function parseQuery(
  query: Partial<{
    [key: string]: string | string[];
  }>,
) {
  if (!query.ethereumWallet) {
    throw new ApiError(400, 'missing ethereumWallet in query');
  }
  let ethereumWallet: string | undefined = undefined;
  try {
    ethereumWallet = ethers.getAddress(query.ethereumWallet.toString());
  } catch (error) {
    throw new ApiError(400, `${query.ethereumWallet} is not an valid ethereum address`);
  }
  if (!ethereumWallet) {
    throw new ApiError(400, `Unknown error while parsing ${query.ethereumWallet} as an ethereum address`);
  }

  return { ethereumWallet };
}
