import re
import sys


def gen(header_file):
  with open(header_file) as f:
    for l in f:
      if not l.startswith('remote_'):
        continue
      parts = re.split(r'\(|\)', l.strip())

      if len(parts) != 3:
        continue

      func_name = parts[0].replace('remote_', '').title().replace('_', '')

      params = []
      for param in parts[1].split(','):
        param_parts = param.split()

        if param.strip().startswith('const'):
          params.append(f'{param_parts[-1].replace("*", "").replace("string", "data")} string')
        else:
          params.append(f'{param_parts[1]} int32')

      print(f'//export {func_name}')
      print(f'func {func_name}({", ".join(params)}) {{')
      print('}')
      print('')

if __name__ == '__main__':
  gen(sys.argv[1])
