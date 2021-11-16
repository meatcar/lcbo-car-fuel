--explain
select
    p.*, s.*,
    p.productTotalVolume / 100 * p.alcoholPercentage as totalAlcoholVolume,
    p.priceInt / totalAlcoholVolume as alcoholCost
from
    stores as s
    left join inventory as i on s.locationNumber = i.locationNumber
    left join products as p on i.product__id = p.itemNumber
where
    upper(s.locationCityName) like :match
limit 10

