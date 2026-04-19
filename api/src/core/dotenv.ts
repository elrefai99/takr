import * as dotenv from 'dotenv';
import * as path from 'path';

const pathEnv = process.env.NODE_ENV === 'production' ? '.env' : '.env.dev';

export default dotenv.config({
     path: path.resolve(process.cwd(), pathEnv),
})
