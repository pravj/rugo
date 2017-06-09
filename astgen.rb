require 'parser/current'

code = File.read('syntax.rb')
parsed_code = Parser::CurrentRuby.parse(code)

p parsed_code
