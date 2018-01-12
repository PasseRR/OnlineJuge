-- 方法一
UPDATE salary
SET sex = (
  CASE
  WHEN sex = 'f'
    THEN 'm'
  WHEN sex = 'm'
    THEN 'f'
  END
);

-- 方法二
UPDATE salary
SET sex = IF(sex = 'f', 'm', 'f');

-- 方法三
UPDATE salary
SET sex = CHAR(ASCII('f') ^ ASCII('m') ^ ASCII(sex));